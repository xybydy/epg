package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/xybydy/epg/golib/mongo"
)

func Save(w http.ResponseWriter, r *http.Request) {
	response := new(apiResponse)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Method == http.MethodPost {
		body, err := r.GetBody()
		if err != nil {
			mes := []byte(err.Error())
			json.Unmarshal(mes, response)
			response.StatusCode = http.StatusInternalServerError
			json.NewEncoder(w).Encode(response)
			return
		}
		defer body.Close()
		res, err := ioutil.ReadAll(body)
		if err != nil {
			mes := []byte(err.Error())
			json.Unmarshal(mes, response)
			response.StatusCode = http.StatusInternalServerError
			json.NewEncoder(w).Encode(response)
			return
		}

		err = mongo.InsertData(res)
		if err != nil {
			mes := []byte(err.Error())
			json.Unmarshal(mes, response)
			response.StatusCode = http.StatusInternalServerError
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
