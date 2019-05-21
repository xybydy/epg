package m3u

type ExtInf struct {
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
