package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/xybydy/epg/golib/mongo"
)

type resp struct {
	ID       interface{} `bson:"_id,omitempty" json:"_id,omitempty"`
	ChanName string      `bson:"chan_name" json:"chan_name,omitempty"`
	TvgID    string      `bson:"tvg_id" json:"tvg_id,omitempty"`
}

// func (r *resp) UnmarshalJSON(data []byte) error {
// 	var s resp

// 	if err := json.Unmarshal(data, &s); err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	return nil
// }

type resps []resp

func Matcher(w http.ResponseWriter, r *http.Request) {

	response := new(apiResponse)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Method == http.MethodPost {
		respData := resps{}
		returnData := resps{}

		res, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(res, &respData)
		if err != nil {
			mes := []byte(err.Error())
			response.StatusCode = http.StatusInternalServerError
			response.Message = string(mes)
			json.NewEncoder(w).Encode(response)
			return
		}

		r.Body.Close()
		if err != nil {
			mes := []byte(err.Error())
			response.StatusCode = http.StatusInternalServerError
			response.Message = string(mes)
			json.NewEncoder(w).Encode(response)
			return
		}

		match, err := mongo.GetData(true, "exact_id")
		log.Println(match)
		for _, v := range respData {
			tvgID := match.GetTvgID(v.ChanName)
			if tvgID == "" {
				continue
			}
			r := resp{ID: v.ID, TvgID: tvgID}
			returnData = append(returnData, r)
		}

		response = &apiResponse{StatusCode: http.StatusOK, Message: fmt.Sprintf("Matched %d channels!", len(returnData)), Data: returnData}
		json.NewEncoder(w).Encode(response)

	} else {
		log.Println(r.Method)
		response = &apiResponse{StatusCode: http.StatusBadRequest, Message: "Bad Request"}
		json.NewEncoder(w).Encode(response)
	}
}
