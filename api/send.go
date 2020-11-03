package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Send(w http.ResponseWriter, r *http.Request) {
	body, err := r.GetBody()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	res, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Errorf(err.Error())
	}

}
