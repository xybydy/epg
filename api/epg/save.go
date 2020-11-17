package epg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/xybydy/epg/golib/mongo"
)

type apiResponse struct {
	StatusCode int                  `json:"status_code"`
	Message    string               `json:"message"`
	Data       mongo.ChannelMatches `json:"data"`
}

func SaveEPG(w http.ResponseWriter, r *http.Request) {
	response := new(apiResponse)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

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
		err = mongo.InsertData(res)
		if err != nil {
			mes := []byte(err.Error())
			response.StatusCode = http.StatusInternalServerError
			response.Message = string(mes)
			json.NewEncoder(w).Encode(response)
			return
		}

		response = &apiResponse{StatusCode: http.StatusOK, Message: "Data successfully sent!"}
		json.NewEncoder(w).Encode(response)

	} else {
		response = &apiResponse{StatusCode: http.StatusBadRequest, Message: "Bad Request"}
		json.NewEncoder(w).Encode(response)
	}
}
