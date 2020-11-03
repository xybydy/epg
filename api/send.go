package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/xybydy/epg/api/golib/mongo"
)

func Send(w http.ResponseWriter, r *http.Request) {
	body, err := r.GetBody()
	defer body.Close()

	if err != nil {
		fmt.Errorf(err.Error())
	}
	res, err := ioutil.ReadAll(body)
	mongo.InsertData(res)
	w.Write([]byte(mongo.DbName))

	if err != nil {
		fmt.Errorf(err.Error())
	}
}
