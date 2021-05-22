package api

import (
	"io"
	"net/http"

	"github.com/xybydy/epg/golib/mongo"
	"github.com/xybydy/epg/golib/utils"
)

func AddChan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			utils.SendApiResponse(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		res, err := mongo.InsertChannel(body, "map")
		if err != nil {
			utils.SendApiResponse(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		utils.SendApiResponse(w, http.StatusOK, "Data successfully sent!", res.InsertedIDs)
	} else {
		utils.SendApiResponse(w, http.StatusBadRequest, "Bad Request", nil)
	}
}
