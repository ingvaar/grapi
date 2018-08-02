package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func updateLineSQL(w http.ResponseWriter, r *http.Request) {
	jsonmap := jsonToMap(w, r)
	pathVars := mux.Vars(r)
	tabName := pathVars["table"]
	id := pathVars["id"]
	multInsert := false

	statement := fmt.Sprintf("UPDATE %s SET ", tabName)
	for key, value := range jsonmap {
		if multInsert {
			statement = fmt.Sprintf("%s, ", statement)
		}
		statement = fmt.Sprintf("%s%s = %s", statement, key, value)
		multInsert = true
	}
	statement = fmt.Sprintf("%s WHERE id=%s;", statement, id)
	_, err := dbSQL.Exec(statement)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}