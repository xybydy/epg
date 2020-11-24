package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xybydy/epg/golib/mongo"
)

func Get(w http.ResponseWriter, r *http.Request) {
	response := new(apiResponse)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Method == http.MethodGet {
		match, err := mongo.GetData()
		if err != nil {
			mes := []byte(err.Error())
			response.StatusCode = http.StatusInternalServerError
			response.Message = string(mes)
			json.NewEncoder(w).Encode(response)
			fmt.Println("3", string(mes))
			return
		}

		response = &apiResponse{StatusCode: http.StatusOK, Message: "Data successfully fetched!", Data: match}
		json.NewEncoder(w).Encode(response)

	} else {
		fmt.Println(r.Method)
		response = &apiResponse{StatusCode: http.StatusBadRequest, Message: "Bad Request"}
		json.NewEncoder(w).Encode(response)
	}
}
