package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/xybydy/epg/api"
)

const DbName = "epg"
const MongoPass = "1ZaaVagptA9N9gJW"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/save", api.Save)
	r.HandleFunc("/api/update", api.Update)
	r.HandleFunc("/api/get", api.Get)
	r.HandleFunc("/api/match", api.Matcher)
	handler := cors.Default().Handler(r)
	log.Println(http.ListenAndServe(":8000", handler))
}
