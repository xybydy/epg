package main

import (
	"epg/m3u"
)

func main() {

	data, i := m3u.ReadM3UFile("m3u/sample.m3u")
	m3u.ImportM3U(data, i)

}
