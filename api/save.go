package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/xybydy/epg/golib/mongo"
)

type apiResponse struct {
	StatusCode int
	Message    string
}

func Save(w http.ResponseWriter, r *http.Request) {
	response := new(apiResponse)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Method == http.MethodPost {
		body, err := r.GetBody()
		fmt.Println(body)
		if err != nil {
			mes := []byte(err.Error())
			json.Unmarshal(mes, response)
			response.StatusCode = http.StatusInternalServerError
			json.NewEncoder(w).Encode(response)
			fmt.Println("1", response)
			return
		}
		defer body.Close()
		res, err := ioutil.ReadAll(body)
		if err != nil {
			mes := []byte(err.Error())
			json.Unmarshal(mes, response)
			response.StatusCode = http.StatusInternalServerError
			json.NewEncoder(w).Encode(response)
			fmt.Println("2", response)
			return
		}
		fmt.Println("77", string(res))
		fmt.Println("4", res)
		err = mongo.InsertData(res)
		if err != nil {
			mes := []byte(err.Error())
			json.Unmarshal(mes, response)
			response.StatusCode = http.StatusInternalServerError
			json.NewEncoder(w).Encode(response)
			fmt.Println("3", response)
			return
		}

		response = &apiResponse{StatusCode: http.StatusOK, Message: "Data successfully sent!"}
		json.NewEncoder(w).Encode(response)

	} else {
		response = &apiResponse{StatusCode: http.StatusBadRequest, Message: "Bad Request"}
		json.NewEncoder(w).Encode(response)
	}
}
