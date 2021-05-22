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

type resps mongo.TvgMaps

func Matcher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Method == http.MethodPost {
		respData := resps{}
		returnData := resps{}

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

		for _, id := range respData {
			if id.Country != "" {
				returnData = append(returnData, mongo.MatchID(id.Name, id.Country, "map"))
			} else {
				returnData = append(returnData, mongo.MatchID(id.Name, "", "map"))
			}
		}

		if len(returnData) == 0 {
			utils.SendApiResponse(w, http.StatusNotFound, "no results found", nil)
		} else if len(returnData) == 1 {
			utils.SendApiResponse(w, http.StatusOK, fmt.Sprintf("Matched %d channels!", len(returnData)), returnData[0])
		} else {
			log.Print(returnData)
			utils.SendApiResponse(w, http.StatusOK, fmt.Sprintf("Matched %d channels!", len(returnData)), returnData)
		}

	} else {
		utils.SendApiResponse(w, http.StatusBadRequest, "Bad Request", nil)
	}
}
