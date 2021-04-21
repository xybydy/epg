const ua_agent = 'f-xtream-api'

import fetch from 'node-fetch'

class XClient {
  constructor(username, password, base_url) {
    this.username = username.trim()
    this.password = password.trim()
    this.ua_agent = ''

    let proper_url = ['http://', 'https://'].some((el) => base_url.trim().startsWith(el))

    if (!proper_url) {
      throw 'URL is not valid. Please correct it'
    } else {
      this.base_url = base_url.trim()
    }

    this.send_request().catch((error) => console.error(`error sending auth request: ${error}`))

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
      console.error(`${wanted_format} is not an allowed output format`)
    }

    if (!(stream_id in this.streams)) {
      throw `${stream_id} is not a valid stream id`
    }

    let stream = this.streams[stream_id]

    return `${this.base_url}/${stream.type}/${this.username}/${this.pasword}/${stream_id}/${wanted_format}`
  }

get_live_categories() {
    let res = this.get_categories('live')

    return res
  }

  get_vod_categories() {
    return this.get_categories('vod')
  }

  get_series_categories() {
    return this.get_categories('series')
  }

  async get_categories(cat_type) {
    let result = ''

    try {
    result = await this.send_request(`get_${cat_type}_categories`)
    } catch (error) {
      console.error(error)
    }
    
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

  async get_streams(stream_action, cat_id) {
    let params = ''
    let results = ''

    if (stream_action !== 'series') {
      stream_action = `${stream_action}_streams`
    }

    if (cat_id !== '') {
      params = `category_id=${cat_id}`
    }

    try {
    results=await this.send_request(`get_${stream_action}`, params)  
    } catch (error) {
      console.error(error)
    }

    for (let item of results) {
      this.streams[item.stream_id] = item
    }

    return results
  }

  async get_series(cat_id) {
    let params = ''
    let results = ''

    if (cat_id !== '') {
      params = `category_id=${cat_id}`
    }

    try {
    results=await this.send_request('get_series', params)  
    } catch (error) {
      console.error(error)
    }

    return results
  }

  async get_series_info(id) {
    let results = ''
    if (id === '') {
      throw 'series id cannot be empty'
    }

    let params = `series_id=${id}`

    try {
    results=await this.send_request('get_series_info', params)  
    } catch (error) {
      console.error(error)
    }

    return results
  }

  async get_vod_info(id) {
    let results = ''

    if (id === '') {
      throw 'vod id cannot be empty'
    }

    let params = `vod_id=${id}`

    try {
    results=await this.send_request('get_vod_info', params)  
    } catch (error) {
      console.error(error)
    }
    
    return results
  }

  get_short_epg(id, limit) {
    return this.get_epg('get_short_epg', id, limit)
  }

  get_long_epg(id) {
    return this.get_epg('get_simple_data_table', id)
  }

  async get_xml() {
    let results = ''
    try {
    results=await this.send_request('xmltv.php')  
    } catch (error) {
      console.error(error)
    }
    

    return results
  }

  async get_epg(action, id, limit) {
    let params = ''
    let results = ''

    if (id ===undefined) {
      throw 'stream id cannot be empty'
    }

    params = `stream_id=${id}`

    if (limit > 0) {
      params = `${params}&limit=${limit}`
    }


    try {
    results= await this.send_request(action, params)  
    } catch (error) {
      console.error(error)
    }
    

    return results
  }

  async send_request(action, params) {
    let result = ''
    let file = 'player_api.php'

    let request_url = `${this.base_url}/${file}?username=${this.username}&password=${this.password}`

    if (action === 'xmltv.php') {
      request_url = `${this.base_url}/${action}?username=${this.username}&password=${this.password}`
    } else if (action !== undefined) {
      request_url = `${this.base_url}/${file}?username=${this.username}&password=${this.password}&action=${action}`
    }

    if (params !== undefined) {
      request_url = `${request_url}&` + encodeURI(params)
    }

    // eslint-disable-next-line no-useless-catch
    try {
      let response = await fetch(request_url, {
        method: 'GET',
        mode: 'cors',
        headers: {
          'User-Agent': this.ua_agent,
        },
      })
      if (action ==='xmltv.php'){
      result = await response.text()
      }else{
      result = await response.json()
      }

      if (action===undefined &&params===undefined&&result.user_info.auth === 0) {
        throw 'user is not authorized'
      }

    } catch (error) {
      throw error
    }


    return result
  }
}

export default XClient
