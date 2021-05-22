package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TvgMaps []TvgMap

// Burada sadece kanal ve tvg_id olacak. kanal ismi unique olmak zorunda ve matcher buradan ilerleyecek
// Collection adi: 	tvg > map
type TvgMap struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name    string             `bson:"name" json:"name"`
	TvgID   string             `bson:"tvg_id" json:"tvg_id"`
	Logo    string             `bson:"logo" json:"logo"`
	Country string             `bson:"country" json:"country"`
}

func (c TvgMaps) GetTvgID(name string) string {
	// name = normalize(name)

	res := BinarySearch(c, name)
	if res == -1 {
		return ""
	}

	return c[res].TvgID
}

func BinarySearch(a TvgMaps, search string) (result int) {
	mid := len(a) / 2

	switch {
	case len(a) == 0:
		result = -1 // not found
	case a[mid].Name > search:
		result = BinarySearch(a[:mid], search)
	case a[mid].Name < search:
		result = BinarySearch(a[mid+1:], search)
		result += mid + 1
	default: // a[mid] == search
		result = mid // found
	}
	return
}
