package api

import (
	"fmt"
	"golib/mongo"
	"io/ioutil"
	"net/http"
)

func Send(w http.ResponseWriter, r *http.Request) {
	body, err := r.GetBody()
	defer body.Close()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	res, err := ioutil.ReadAll(body)
	mongo.InsertData(res)
	w.Write(mongo.DbName)
	if err != nil {
		fmt.Errorf(err.Error())
	}
}
