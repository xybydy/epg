package mongo

type TvgIds []TvgId

type TvgId struct {
	ChanName string `bson:"chan_name" json:"chan_name"`
	TvgId    string `bson:"tvg_id" json:"tvg_id"`
}
