package main

import (
	"strconv"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func status(w http.ResponseWriter, r *http.Request) {
	err := db.Ping()

	if err == nil {
		fmt.Fprintln(w, "Database connected")
	} else {
		fmt.Fprintln(w, err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome !")
}

func createLine(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented")
}

func deleteLine(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	tab_name := pathVars["table"]
	id := pathVars["id"]
	id_num, err_atoi := strconv.Atoi(id)

	statement := fmt.Sprintf("DELETE FROM %s WHERE id=%d", tab_name, id_num)
	if err_atoi != nil {
		fmt.Fprintf(w, "Error: invalid id '%s'\n", id)
	} else {
		_, err := db.Query(statement)
		if err != nil {
			fmt.Fprintf(w, "%s", err)
		} else {
			fmt.Fprintf(w, "Line %s deleted", id)
		}
	}
}
