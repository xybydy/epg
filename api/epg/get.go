package epg

import (
	"encoding/json"
	"net/http"

	"github.com/xybydy/epg/golib/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func GetEPG(w http.ResponseWriter, r *http.Request) {
	response := new(apiResponse)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodGet {
		data, err := mongo.GetData(bson.D{{}})
		if err != nil {
			response = &apiResponse{StatusCode: http.StatusBadRequest, Message: err.Error(), Data: data}
			json.NewEncoder(w).Encode(response)
			return
		}

		response = &apiResponse{StatusCode: http.StatusOK, Message: "Data successfully sent!", Data: data}
		json.NewEncoder(w).Encode(response)

	} else {
		response = &apiResponse{StatusCode: http.StatusBadRequest, Message: "Bad Request"}
		json.NewEncoder(w).Encode(response)
	}
}
