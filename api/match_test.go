package api

import (
	"log"
	"testing"

	"github.com/xybydy/epg/golib/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func TestMatcher(t *testing.T) {
	var c []interface{}

	match, err := mongo.GetData(true, "tvs")
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	for _, m := range match {
		if m.TvgID != "" {
			c = append(c, bson.M{"chan_name": m.ChanName, "tvg_id": m.TvgID})
		}
	}

	if mongo.InsertData(c, "exact_id") != nil {
		log.Println(err)
		t.Fail()
	}
}
