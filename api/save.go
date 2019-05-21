package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/xybydy/epg/golib/mongo"
)

type apiResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func Save(w http.ResponseWriter, r *http.Request) {
	response := new(apiResponse)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Method == http.MethodPost {
		res, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			mes := []byte(err.Error())
			response.StatusCode = http.StatusInternalServerError
			response.Message = string(mes)
			json.NewEncoder(w).Encode(response)
			return
		}
		err = mongo.InsertChannels(res, "tvs")
		if err != nil {
			mes := []byte(err.Error())
			response.StatusCode = http.StatusInternalServerError
			response.Message = string(mes)
			json.NewEncoder(w).Encode(response)
			log.Println("3", string(mes))
			return
		}

		response = &apiResponse{StatusCode: http.StatusOK, Message: "Data successfully sent!"}
		json.NewEncoder(w).Encode(response)

	} else {
		log.Println(r.Method)
		response = &apiResponse{StatusCode: http.StatusBadRequest, Message: "Bad Request"}
		json.NewEncoder(w).Encode(response)
	}
}
