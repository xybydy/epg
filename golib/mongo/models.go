package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChannelMatches []ChannelMatch

type ChannelMatch struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	// ChanID     int                `bson:"id" json:"id"`
	ChanName   string `bson:"chan_name" json:"chan_name"`
	TvgID      string `bson:"tvg_id" json:"tvg_id"`
	UserName   string `bson:"user_name" json:"user_name"`
	GroupTitle string `bson:"group_title" json:"group_title"`
}

func (c ChannelMatches) GetTvgID(name string) string {
	// name = normalize(name)

	res := BinarySearch(c, name)
	if res == -1 {
		return ""
	}

	return c[res].TvgID
}

func BinarySearch(a ChannelMatches, search string) (result int) {
	mid := len(a) / 2

	switch {
	case len(a) == 0:
		result = -1 // not found
	case a[mid].ChanName > search:
		result = BinarySearch(a[:mid], search)
	case a[mid].ChanName < search:
		result = BinarySearch(a[mid+1:], search)
		result += mid + 1
	default: // a[mid] == search
		result = mid // found
	}
	return
}
