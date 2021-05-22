package mongo

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestInsertData(t *testing.T) {
	var a map[string][]TvgMap
	var b []interface{}

	f, err := os.ReadFile("aa.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	json.Unmarshal(f, &a)

	for _, v := range a {
		for _, i := range v {
			b = append(b, i)
		}
	}

	InsertData(b, "map")

}
