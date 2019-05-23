package m3u

import (
	"bufio"
	"os"
	"strings"

	"github.com/xybydy/m3u8"
)

type M3U8Playlist []*ExtInf

type ExtInf struct {
	SeqId        uint64
	Duration     float64
	Title        string
	GroupTitle   string `json:"group-title,omitempty"`
	ParentCode   string `json:"parent-code,omitempty"`
	AudioTrack   string `json:"audio-track,omitempty"`
	TvgId        string `json:"tvg-id,omitempty"`
	TvgLogo      string `json:"tvg-logo,omitempty"`
	TvgLogoSmall string `json:"tvg-logo_small,omitempty"`
	TvgShift     string `json:"tvg-shift,omitempty"`
}

var (
	RawM3U8  *m3u8.MediaPlaylist
	Playlist *M3U8Playlist
)

func readM3UFile(path string) (string, uint) {
	buf := make([]string, 0)
	var size uint
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		buf = append(buf, scanner.Text())
		if strings.HasPrefix(scanner.Text(), "#EXTINF") {
			size++
		}
	}

	return strings.Join(buf, "\n"), size

}

func ImportM3U(path string) {
	input, size := readM3UFile(path)

	RawM3U8, err := m3u8.NewMediaPlaylist(8, size)
	if err != nil {
		panic(err)
	}
	err = RawM3U8.DecodeFrom(strings.NewReader(input), false)
	if err != nil {
		panic(err)
	}
}
