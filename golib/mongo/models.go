package mongo

type ChannelMatches []ChannelMatch

type ChannelMatch struct {
	ChanName string `bson:"chan_name" json:"chan_name"`
	TvgID    string `bson:"tvg_id" json:"tvg_id"`
	UserName string ´bson:"user_name" json:"user_name"´
}
