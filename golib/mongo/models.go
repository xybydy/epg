package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChannelMatches []ChannelMatch

type ChannelMatch struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ChanName   string             `bson:"chan_name" json:"chan_name"`
	TvgID      string             `bson:"tvg_id" json:"tvg_id"`
	UserName   string             `bson:"user_name" json:"user_name"`
	GroupTitle string             `bson:"group_title" json:"group_title"`
}
