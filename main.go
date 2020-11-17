package main

import (
	// 	"fmt"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
	api "github.com/xybydy/epg/api/epg"
)

const DbName = "epg"
const MongoPass = "1ZaaVagptA9N9gJW"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/epg/save", api.SaveEPG)
	r.HandleFunc("/api/epg/get", api.GetEPG)
	fmt.Println(http.ListenAndServe(":8000", r))
}

// func main() {
// 	// eh := &ChannelMatches{}
// 	a, err := mongo.GetData(bson.D{{}})
// 	if err != nil {
// 		fmt.Println("qq", err)
// 	}
// 	fmt.Println(a[0].ChanName)
// 	// bson.UnmarshalJSON()
// }
