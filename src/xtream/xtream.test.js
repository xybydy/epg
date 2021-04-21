import auth from './auth'
import XClient from './xtream'

let xc

beforeAll(async () => {
      xc = new XClient(auth.username, auth.password, auth.url)
      await xc.init_data()
      })


test('get_stream_url',  async () => {
await xc.get_live_stream()
    // {
    //     category_id: '5',
    //     category_name: 'TR ➾ BELGESEL',
    //     parent_id: 0,
    //     type: 'live'
    //   }
console.log(xc.get_stream_url(15900))

})

test('get_live_cats',  async () => {
    // {
    //     category_id: '5',
    //     category_name: 'TR ➾ BELGESEL',
    //     parent_id: 0,
    //     type: 'live'
    //   }
   await xc.get_live_categories()

})

test('get_vod_cats', async () => {
      //  {
      //   category_id: '90',
      //   category_name: 'VOD ➾ English Movies',
      //   parent_id: 0,
      //   type: 'vod'
      // }

await xc.get_vod_categories()
})

test('get_series_cats', async () => {
      //  {
      //   category_id: '179',
      //   category_name: 'DE ➾ Netflix Series',
      //   parent_id: 0,
      //   type: 'series'
      // },

await xc.get_series_categories()
})

test('get_live_stream', async () => {
      //  {
      //   num: 55,
      //   name: 'TR: YABAN TV',
      //   stream_type: 'live',
      //   stream_id: 15900,
      //   stream_icon: 'http://axe.tvlogos.xyz:8080/logos/yabantv.png',
      //   epg_channel_id: 'yabantv.tr',
      //   added: '1578305427',
      //   category_id: '5',
      //   custom_sid: '',
      //   tv_archive: 0,
      //   direct_source: '',
      //   tv_archive_duration: 0
      // },

await xc.get_live_stream(5)
})


test('get_vod_stream', async () => {
      //   {
      //   num: 99,
      //   name: 'Dinosaur Island 2014 (ENG)',
      //   stream_type: 'movie',
      //   stream_id: 19569,
      //   stream_icon: 'https://image.tmdb.org/t/p/w600_and_h900_bestv2/vP0muFZYdlb0GfFZgnoPrgMmuE8.jpg',
      //   rating: '4.6',
      //   rating_5based: 2.3,
      //   added: '1578703970',
      //   category_id: '90',
      //   container_extension: 'mp4',
      //   custom_sid: '',
      //   direct_source: ''
      // },

await xc.get_vod_streams(90)
})

test('get_series', async () => {
      // {
      //   num: 1,
      //   name: 'The Mandalorian',
      //   series_id: 186,
      //   cover: 'https://image.tmdb.org/t/p/w600_and_h900_bestv2/BbNvKCuEF4SRzFXR16aK6ISFtR.jpg',
      //   plot: "İmparatorluğun çöküşünden sonra ve ilk düzenin ortaya çıkmasından önce, galaksinin dış kesimlerinde Yeni Cumhuriyet'in otoritesinden uzak bir şekilde yalnız bir silahlı savaşçının mücadelesini takip ediyoruz.",
      //   cast: 'Pedro Pascal',
      //   director: '',
      //   genre: 'Bilim Kurgu & Fantazi, Aksiyon & Macera',
      //   releaseDate: '2019-11-12',
      //   last_modified: '1615382851',
      //   rating: '8',
      //   rating_5based: 4,
      //   backdrop_path: [
      //     'https://image.tmdb.org/t/p/w1280/o7qi2v4uWQ8bZ1tW3KI0Ztn2epk.jpg'
      //   ],
      //   youtube_trailer: '3bSCegaFRdo',
      //   episode_run_time: '35',
      //   category_id: '249'
      // },

await xc.get_series(249)
})

test('get_series_info', async () => {
// {
//       seasons: [
//         {
//           air_date: '2019-11-12',
//           episode_count: 8,
//           id: 110346,
//           name: '1. Sezon',
//           overview: '',
//           season_number: 1,
//           cover: 'https://image.tmdb.org/t/p/w600_and_h900_bestv2/sUTqIb82LxYhPT0SfI8AR03GLpz.jpg',
//           cover_big: 'https://image.tmdb.org/t/p/w600_and_h900_bestv2/sUTqIb82LxYhPT0SfI8AR03GLpz.jpg'
//         },
//         {
//           air_date: '2020-10-30',
//           episode_count: 8,
//           id: 153254,
//           name: '2. Sezon',
//           overview: '',
//           season_number: 2,
//           cover: 'https://image.tmdb.org/t/p/w600_and_h900_bestv2/b8cs9DzhxRQLQ7TvfLihYbAC6fg.jpg',
//           cover_big: 'https://image.tmdb.org/t/p/w600_and_h900_bestv2/b8cs9DzhxRQLQ7TvfLihYbAC6fg.jpg'
//         }
//       ],
//       info: {
//         name: 'The Mandalorian',
//         cover: 'https://image.tmdb.org/t/p/w600_and_h900_bestv2/BbNvKCuEF4SRzFXR16aK6ISFtR.jpg',
//         plot: "İmparatorluğun çöküşünden sonra ve ilk düzenin ortaya çıkmasından önce, galaksinin dış kesimlerinde Yeni Cumhuriyet'in otoritesinden uzak bir şekilde yalnız bir silahlı savaşçının mücadelesini takip ediyoruz.",
//         cast: 'Pedro Pascal',
//         director: '',
//         genre: 'Bilim Kurgu & Fantazi, Aksiyon & Macera',
//         releaseDate: '2019-11-12',
//         last_modified: '1615382851',
//         rating: '8',
//         rating_5based: 4,
//         backdrop_path: [
//           'https://image.tmdb.org/t/p/w1280/o7qi2v4uWQ8bZ1tW3KI0Ztn2epk.jpg'
//         ],
//         youtube_trailer: '3bSCegaFRdo',
//         episode_run_time: '35',
//         category_id: '249'
//       },
//       episodes: {
//         '1': [
//           [Object], [Object],
//           [Object], [Object],
//           [Object], [Object],
//           [Object], [Object]
//         ],
//         '2': [
//           [Object], [Object],
//           [Object], [Object],
//           [Object], [Object],
//           [Object], [Object]
//         ]
//       }
//     }
await xc.get_series_info(186)
})

test('get_series', async () => {
      // {
      //   num: 1,
      //   name: 'The Mandalorian',
      //   series_id: 186,
      //   cover: 'https://image.tmdb.org/t/p/w600_and_h900_bestv2/BbNvKCuEF4SRzFXR16aK6ISFtR.jpg',
      //   plot: "İmparatorluğun çöküşünden sonra ve ilk düzenin ortaya çıkmasından önce, galaksinin dış kesimlerinde Yeni Cumhuriyet'in otoritesinden uzak bir şekilde yalnız bir silahlı savaşçının mücadelesini takip ediyoruz.",
      //   cast: 'Pedro Pascal',
      //   director: '',
      //   genre: 'Bilim Kurgu & Fantazi, Aksiyon & Macera',
      //   releaseDate: '2019-11-12',
      //   last_modified: '1615382851',
      //   rating: '8',
      //   rating_5based: 4,
      //   backdrop_path: [
      //     'https://image.tmdb.org/t/p/w1280/o7qi2v4uWQ8bZ1tW3KI0Ztn2epk.jpg'
      //   ],
      //   youtube_trailer: '3bSCegaFRdo',
      //   episode_run_time: '35',
      //   category_id: '249'
      // },

await xc.get_series(249)
})

test('get_vod_info', async () => {
//  {
//       info: {
//         kinopoisk_url: 'https://www.themoviedb.org/movie/294093',
//         tmdb_id: '294093',
//         name: 'Dinosaur Island 2014 (ENG)',
//         o_name: 'Dinosaur Island 2014 (ENG)',
//         cover_big: 'https://image.tmdb.org/t/p/w600_and_h900_bestv2/vP0muFZYdlb0GfFZgnoPrgMmuE8.jpg',
//         movie_image: 'https://image.tmdb.org/t/p/w600_and_h900_bestv2/vP0muFZYdlb0GfFZgnoPrgMmuE8.jpg',
//         releasedate: '2014-09-01',
//         episode_run_time: '90',
//         youtube_trailer: 'U18sOypRW8s',
//         director: 'Matt Drummond',
//         actors: 'Darius Williams-Watt, Kate Rasmussen, Joe Bistaveous, Juliette Frederick, Vincent Naviti',
//         cast: 'Darius Williams-Watt, Kate Rasmussen, Joe Bistaveous, Juliette Frederick, Vincent Naviti',
//         description: 'The adventure begins when Lucas, a 13 year old boy, embarks on the vacation of a lifetime. When disaster strikes, Lucas finds himself stranded in a strange land littered with ghost ships and prehistoric creatures. While searching for other signs of life, Lucas hears a radio broadcast in the distance and is drawn into the jungle where he encounters a beautiful young girl who claims to have come from the 1950s. Together they set out on a quest to get home all the while uncovering secrets that will forever change the future.',
//         plot: 'The adventure begins when Lucas, a 13 year old boy, embarks on the vacation of a lifetime. When disaster strikes, Lucas finds himself stranded in a strange land littered with ghost ships and prehistoric creatures. While searching for other signs of life, Lucas hears a radio broadcast in the distance and is drawn into the jungle where he encounters a beautiful young girl who claims to have come from the 1950s. Together they set out on a quest to get home all the while uncovering secrets that will forever change the future.',
//         age: '',
//         mpaa_rating: '',
//         rating_count_kinopoisk: 0,
//         country: 'Australia',
//         genre: 'Adventure, Science Fiction, Fantasy',
//         backdrop_path: [
//           'https://image.tmdb.org/t/p/w1280/4iJ43ltaXM52tcjnOSIF21dXMLE.jpg'
//         ],
//         duration_secs: 4981,
//         duration: '01:23:01',
//         video: {
//           index: 0,
//           codec_name: 'h264',
//           codec_long_name: 'H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10',
//           profile: 'Main',
//           codec_type: 'video',
//           codec_time_base: '39855817/1911168000',
//           codec_tag_string: 'avc1',
//           codec_tag: '0x31637661',
//           width: 1280,
//           height: 720,
//           coded_width: 1280,
//           coded_height: 720,
//           has_b_frames: 2,
//           sample_aspect_ratio: '1:1',
//           display_aspect_ratio: '16:9',
//           pix_fmt: 'yuv420p',
//           level: 40,
//           color_range: 'tv',
//           color_space: 'bt709',
//           color_transfer: 'bt709',
//           color_primaries: 'bt709',
//           chroma_location: 'left',
//           refs: 1,
//           is_avc: 'true',
//           nal_length_size: '4',
//           r_frame_rate: '24000/1001',
//           avg_frame_rate: '955584000/39855817',
//           time_base: '1/16000',
//           start_pts: 0,
//           start_time: '0.000000',
//           duration_ts: 79711634,
//           duration: '4981.977125',
//           bit_rate: '875222',
//           bits_per_raw_sample: '8',
//           nb_frames: '119448',
//           disposition: [Object],
//           tags: [Object]
//         },
//         audio: {
//           index: 1,
//           codec_name: 'aac',
//           codec_long_name: 'AAC (Advanced Audio Coding)',
//           profile: 'LC',
//           codec_type: 'audio',
//           codec_time_base: '1/48000',
//           codec_tag_string: 'mp4a',
//           codec_tag: '0x6134706d',
//           sample_fmt: 'fltp',
//           sample_rate: '48000',
//           channels: 2,
//           channel_layout: 'stereo',
//           bits_per_sample: 0,
//           r_frame_rate: '0/0',
//           avg_frame_rate: '0/0',
//           time_base: '1/48000',
//           start_pts: 0,
//           start_time: '0.000000',
//           duration_ts: 239135720,
//           duration: '4981.994167',
//           bit_rate: '159419',
//           max_bit_rate: '159419',
//           nb_frames: '233531',
//           disposition: [Object],
//           tags: [Object]
//         },
//         bitrate: 1041,
//         rating: '4.6'
//       },
//       movie_data: {
//         stream_id: 19569,
//         name: 'Dinosaur Island 2014 (ENG)',
//         added: '1578703970',
//         category_id: '90',
//         container_extension: 'mp4',
//         custom_sid: '',
//         direct_source: ''
//       }
//     }
await xc.get_vod_info(19569)
})

test('get_short_epg', async () => {
    // {
    //   epg_listings: [
    //     {
    //       id: '39991367',
    //       epg_id: '8',
    //       title: 'TW9yZ2FuIEZyZWVtYW4gxLBsZS4uLg==',
    //       lang: '',
    //       start: '2021-04-21 23:30:00',
    //       end: '2021-04-22 00:20:00',
    //       description: 'QmlsaW7Dp2FsdMSxbsSxbiBzxLFybGFyxLEgYmlsaW0gaW5zYW5sYXLEsSwgaWxrZWwga29ya3UgdmUgYXJ6dWxhcsSxbcSxesSxbiBrYXluYcSfxLEgb2xhcmFrIGJpbGluZW4gYmlsaW7Dp2FsdMSxbsSxbiBiaXppIG5hc8SxbCBrb3J1eWFiaWxlY2XEn2luaSB2ZSBoYXN0YWzEsWtsYXJhIGthcsWfxLEgdsO8Y3VkdW11enUgaXlpbGXFn3RpcmViaWxlY2XEn2luaSBrZcWfZmV0dGku',
    //       channel_id: 'discoveryscience.tr',
    //       start_timestamp: '1619037000',
    //       stop_timestamp: '1619040000'
    //     },
    //     {
    //       id: '39991368',
    //       epg_id: '8',
    //       title: 'TW9yZ2FuIEZyZWVtYW4gxLBsZS4uLg==',
    //       lang: '',
    //       start: '2021-04-22 00:20:00',
    //       end: '2021-04-22 01:05:00',
    //       description: 'RWJlZGl5ZXQgc29uYSBlcmVjZWsgbWk/IFphbWFuIHNvbnN1emEgZGVrIHPDvHJlYmlsaXIgbWk/IEViZWRpeWV0IHphdGVuIHZhciBvbGFiaWxpciB2ZSBnZWxlY2VrIMWfaW1kaXlpIMWfZWtpbGxlbmRpcm1layBpw6dpbiBnZcOnbWnFn2UgeWFwxLFsYW4gYmlyIHlvbGN1bHVrIG9sYWJpbGlyIHZleWEgemFtYW4gaG9sb2dyYWZpayBiaXIgaXpkw7zFn8O8bSBvbGFiaWxpci4=',
    //       channel_id: 'discoveryscience.tr',
    //       start_timestamp: '1619040000',
    //       stop_timestamp: '1619042700'
    //     },
    //     {
    //       id: '39991369',
    //       epg_id: '8',
    //       title: 'QmlsaW1pbiBTaWhyaQ==',
    //       lang: '',
    //       start: '2021-04-22 01:05:00',
    //       end: '2021-04-22 01:30:00',
    //       description: 'S2F6xLFsxLEgecSxa8SxbSBpemxleWljaWxlcmkgxZ9va2UgZWRlbiB5ZW5pIG51bWFyYWxhcmxhIGVrcmFuYSBkw7ZuZW4gZGl6aWRlIGTDvG55YW7EsW4gZW4gw7xubMO8IHNpaGlyYmF6bGFyxLEgYmlsaW1pIGlsbMO8enlvbmEgZMO2bsO8xZ90w7xyw7x5b3IgdmUgaGVyIG51bWFyYW7EsW4gYXJkxLFuZGFraSBiaWxpbXNlbCBnZXLDp2XEn2kgYcOnxLFrbMSxeW9ybGFyLg==',
    //       channel_id: 'discoveryscience.tr',
    //       start_timestamp: '1619042700',
    //       stop_timestamp: '1619044200'
    //     },
    //     {
    //       id: '39991370',
    //       epg_id: '8',
    //       title: 'QmlsaW1pbiBTaWhyaQ==',
    //       lang: '',
    //       start: '2021-04-22 01:30:00',
    //       end: '2021-04-22 01:50:00',
    //       description: 'Qm/FnyBrYWxtxLHFnyBiaW5hIGl6bGV5aWNpbGVyaSDFn29rZSBlZGVuIHllbmkgbnVtYXJhbGFybGEgZWtyYW5hIGTDtm5lbiBkaXppZGUgZMO8bnlhbsSxbiBlbiDDvG5sw7wgc2loaXJiYXpsYXLEsSBiaWxpbWkgaWxsw7x6eW9uYSBkw7Zuw7zFn3TDvHLDvHlvciB2ZSBoZXIgbnVtYXJhbsSxbiBhcmTEsW5kYWtpIGJpbGltc2VsIGdlcsOnZcSfaSBhw6fEsWtsxLF5b3JsYXIu',
    //       channel_id: 'discoveryscience.tr',
    //       start_timestamp: '1619044200',
    //       stop_timestamp: '1619045400'
    //     }
    //   ]
    // }
await xc.get_short_epg(2053)
})

test('get_long_epg', async () => {
    // [ {
    //       id: '38711763',
    //       epg_id: '8',
    //       title: 'QsO8ecO8ayBCZXlpbiBLdXJhbcSx',
    //       lang: '',
    //       start: '2021-04-15 01:30:00',
    //       end: '2021-04-15 02:15:00',
    //       description: 'SGVzYXB0YSBvbG1heWFuLiBZYXLEscWfbWFjxLFsYXIgaW1rYW5zxLF6IGdpYmkgZ8O2csO8bmVuIGJpciBtw7xoZW5kaXNsaWsgcHJvamVzaW5pIGdlcsOnZWtsZcWfdGlybWVrIHpvcnVuZGFsYXIuIFpla2FsYXLEsW7EsSBrdWxsYW5hcmFrIGJpciDDp8O2esO8bSDDvHJldG1layBpw6dpbiBzYWRlY2UgMzAgZGFraWthbGlrIGJpciB6YW1hbmxhcsSxIHZhci4=',
    //       channel_id: 'discoveryscience.tr',
    //       start_timestamp: '1618439400',
    //       stop_timestamp: '1618442100',
    //       now_playing: 0,
    //       has_archive: 0
    //     },
    //     {
    //       id: '38711764',
    //       epg_id: '8',
    //       title: 'QsO8ecO8ayBCZXlpbiBLdXJhbcSx',
    //       lang: '',
    //       start: '2021-04-15 02:15:00',
    //       end: '2021-04-15 03:00:00',
    //       description: 'QXJhYmFsYXLEsSB5YWthbGFtYWsuIFlhcsSxxZ9tYWPEsWxhciDDp8O2esO8bcO8IGlta2Fuc8SxeiBnaWJpIGfDtnLDvG5lbiBiaXIgbcO8aGVuZGlzbGlrIHByb2JsZW1pbmkgw6fDtnptZXllIMOnYWzEscWfYWNha2xhci4gQmV5aW4gZ8O8w6dsZXJpbmkga3VsbGFuYXJhayBiaXIgw6fDtnrDvG0gw7xyZXRlYmlsbWVrIGnDp2luIHNhZGVjZSAzMCBkYWtpa2FsYXLEsSB2YXIu',
    //       channel_id: 'discoveryscience.tr',
    //       start_timestamp: '1618442100',
    //       stop_timestamp: '1618444800',
    //       now_playing: 0,
    //       has_archive: 0
    //     },]
await xc.get_long_epg(2053)
})
