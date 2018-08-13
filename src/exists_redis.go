package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// existsRedis : do an exists on the id passed in the url an http response code
func existsRedis(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	id := fmt.Sprintf("%s:%s", pathVars["type"], pathVars["id"])

	reply, err := redisCli.Cmd("EXISTS", id).Int()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else if reply == 0 {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
