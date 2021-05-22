package api

import (
	"log"
	"testing"

	"github.com/xybydy/epg/golib/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func TestMatcher(t *testing.T) {
	var c []interface{}

	match, err := mongo.GetData(true, "map")
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	for _, m := range match {
		if m.TvgID != "" {
			c = append(c, bson.M{"chan_name": m.Name, "tvg_id": m.TvgID})
		}
	}

	_, err = mongo.InsertData(c, "exact_id")
	if err != nil {
		log.Println(err)
		t.Fail()
	}
}
