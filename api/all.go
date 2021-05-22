package api

import (
	"net/http"

	"github.com/xybydy/epg/golib/mongo"
	"github.com/xybydy/epg/golib/utils"
)

func All(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Method == http.MethodGet {
		match, err := mongo.GetData(false, "map")
		if err != nil {
			utils.SendApiResponse(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		utils.SendApiResponse(w, http.StatusOK, "Data successfully fetched!", match)
	} else {
		utils.SendApiResponse(w, http.StatusBadRequest, "Bad Request", nil)
	}
}
