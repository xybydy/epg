const ua_agent = 'f-xtream-api'

class XClient {
  constructor(username, password, base_url) {
    this.username = username.trim()
    this.password = password.trim()
    this.ua_agent = ''

    let proper_url = [('http://', 'https://')].some((el) => base_url.trim().startsWith(el))

    if (!proper_url) {
      throw 'URL is not valid. Please correct it'
    } else {
      this.base_url = base_url.trim()
    }

    let auth_data = ''
    this.send_request()
      .then((resp) => (auth_data = resp))
      .catch(console.error((e) => 'error sending auth request: ${e}'))

    console.log(auth_data)

    this.server_info = ''
    this.user_info = ''

    this.streams = {}
  }

  get_stream_url(stream_id, wanted_format) {
    let valid_format = false

    for (let element of this.user_info['user_info']['allowed_output_formats']) {
      if (wanted_format === element) {
        valid_format = true
        break
      }
    }

    if (!valid_format) {
      console.error('${wanted_format} is not an allowed output format')
    }

    if (!(stream_id in this.streams)) {
      throw '${stream_id} is not a valid stream id'
    }

    let stream = this.streams[stream_id]

    return '${this.base_url}/${stream.type}/${this.username}/${this.pasword}/${stream_id}/${wanted_format}'
  }

  get_live_categories() {
    return this.get_categories('live')
  }

  get_vod_categories() {
    return this.get_categories('vod')
  }

  get_series_categories() {
    return this.get_categories('series')
  }

  get_categories(cat_type) {
    let result = ''
    this.send_request('get_${cat_type}_categories')
      .then((resp) => (result = resp))
      .catch((error) => console.error(error))

    for (let item of result) {
      item.type = cat_type
    }
    return result
  }

  get_live_stream(cat_id) {
    return this.get_streams('live', cat_id)
  }

  get_vod_streams(cat_id) {
    return this.get_streams('vod', cat_id)
  }

  get_streams(stream_action, cat_id) {
    let params = ''
    let results = ''

    if (stream_action !== 'series') {
      stream_action = '${stream_action}_streams'
    }

    if (cat_id !== '') {
      params = 'category_id=${cat_id}'
    }

    this.send_request('get_${stream_action}', params)
      .then((resp) => (results = resp))
      .catch((error) => console.error(error))

    for (let item of results) {
      this.streams[item.stream_id] = item
    }

    return results
  }

  get_series(cat_id) {
    let params = ''
    let results = ''

    if (cat_id !== '') {
      params = 'category_id=${cat_id}'
    }

    this.send_request('get_series', params)
      .then((resp) => (results = resp))
      .catch((error) => console.error(error))

    return results
  }

  get_series_info(id) {
    let results = ''
    if (id === '') {
      throw 'series id cannot be empty'
    }

    let params = 'series_id=${id}'

    this.send_request('get_series_info', params)
      .then((resp) => (results = resp))
      .catch((error) => console.error(error))

    return results
  }

  get_vod_info(id) {
    let results = ''

    if (id === '') {
      throw 'vod id cannot be empty'
    }

    let params = 'vod_id=${id}'

    this.send_request('get_vod_info', params)
      .then((resp) => (results = resp))
      .catch((error) => console.error(error))
  }

  get_short_epg(id, limit) {
    return this.get_epg('get_short_epg', id, limit)
  }

  get_long_epg(id) {
    return this.get_epg('get_simple_data_table', id)
  }

  get_xml() {
    let results = ''
    this.send_request('xmltv.php')
      .then((resp) => (results = resp))
      .catch((error) => console.log(error))

    return results
  }

  get_epg(action, id, limit) {
    let params = ''
    let results = ''

    if (id !== '') {
      throw 'stream id cannot be empty'
    }

    params = 'stream_id=${id}'

    if (limit > 0) {
      params = '${params}&limit=${limit}'
    }

    this.send_request(action, params)
      .then((resp) => (results = resp))
      .catch((error) => console.error(error))

    return results
  }

  async send_request(action, params) {
    let request_url = ''
    let file = 'player_api.php'
    let result = ''

    if (action === 'xmltv.php') {
      request_url = '${this.base_url}/${file}?username=${this.username}&password=${this.password}'
    } else if (action !== '') {
      request_url =
        '${this.base_url}/${file}?username=${this.username}&password=${this.password}&action=${action}'
    }

    if (params !== '') {
      request_url = '${request_url}&' + encodeURI(params)
    }

    fetch(request_url, {
      method: 'GET',
      mode: 'cors',
      headers: {
        'User-Agent': this.ua_agent,
      },
    })
      .then((resp) => resp.json())
      .then((res) => (result = res))
      .catch((error) => {
        throw error
      })

    return result
  }
}

export default XClient
