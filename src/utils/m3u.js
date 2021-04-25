export function SanitizeM3u(m3uFile) {
  let outPut = []

  for (var [key, item] of Object.entries(m3uFile)) {
    outPut.push({
      _id: key,
      chan_name: item.name,
      url: item.url,
      raw: item.raw,
      tvg_id: item.tvg.id,
      tvg_name: item.tvg.name,
      tvg_language: item.tvg.language,
      tvg_country: item.tvg.country,
      tvg_logo: item.tvg.logo,
      tvg_url: item.tvg.url,
      group_title: item.group.title,
      http_referrer: item.http.referrer,
      http_user_agent: item.http['user-agent'],
    })
  }
  return outPut
}

export function ExportM3u(data, filename = 'epg.m3u8') {
  let out = '#EXTM3U\n'
  for (let item of Object.values(data)) {
    out += `#EXTINF:-1, tvg-id="${item.tvg_id}" tvg-name="${item.tvg_name}" tvg-logo="${item.tvg_logo}" group-title="${item.group_title}",${item.name}\n${item.url}\n`
    // out += item.raw + '\n'
  }

  let blob = new Blob([out], { type: 'text/plain' })
  const blobUrl = URL.createObjectURL(blob)
  const anchor = document.createElement('a')
  anchor.href = blobUrl
  anchor.target = '_blank'
  anchor.download = filename
  anchor.click()
  URL.revokeObjectURL(blobUrl)

  return blob
}

export function BuildUrl(username, password, url) {
  url = url.endsWith('/') ? url : url + '/'

  return `${url}get.php?username=${username}&password=${password}&type=m3u_plus&output=ts`
}
