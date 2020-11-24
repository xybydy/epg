package main

import (
	"fmt"
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
	handler := cors.Default().Handler(r)
	fmt.Println(http.ListenAndServe(":8000", handler))
}
