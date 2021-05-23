package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/xybydy/epg/golib/mongo"
	"github.com/xybydy/epg/golib/utils"
)

func Complete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Method == http.MethodPost {
		respData := struct{ Query string }{}

		res, err := io.ReadAll(r.Body)
		err = json.Unmarshal(res, &respData)
		if err != nil {
			utils.SendApiResponse(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		r.Body.Close()
		if err != nil {
			utils.SendApiResponse(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		returnData := mongo.AutoComplete(respData.Query, "map")

		if len(returnData) == 0 {
			utils.SendApiResponse(w, http.StatusNotFound, "no results found", nil)
		} else {
			log.Print(returnData)
			utils.SendApiResponse(w, http.StatusOK, fmt.Sprintf("Matched %d channels!", len(returnData)), returnData)
		}

	} else {
		utils.SendApiResponse(w, http.StatusBadRequest, "Bad Request", nil)
	}
}
