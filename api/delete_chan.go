package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/xybydy/epg/golib/mongo"
	"github.com/xybydy/epg/golib/utils"
)

func DeleteChan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Method == http.MethodPost {
		res, err := io.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			utils.SendApiResponse(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		s := []string{}
		json.Unmarshal(res, &s)

		err = mongo.DeleteChannels(s, "map")
		if err != nil {
			utils.SendApiResponse(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		utils.SendApiResponse(w, http.StatusOK, "Data successfully deleted!", s)
	} else {
		utils.SendApiResponse(w, http.StatusBadRequest, "Bad Request", nil)
	}
}
