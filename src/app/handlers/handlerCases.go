package handlers

import (
	readjson "assignment_2/src/app/readJson"
	"fmt"
	"net/http"
	"path"
)

/*
Function to handle different types of https requests. switch case that handle GET request spesificly all other are handled under default.*/
func CasesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		casesGetRequest(w, r)
	default:
		fmt.Fprintf(w, "The only http method implemented is the GET request")
	}
}

func casesGetRequest(w http.ResponseWriter, r *http.Request) {
	search := path.Base(r.URL.Path)
	if search != "cases" {

		w.Header().Set("content-type", "application/json")

		fmt.Fprintf(w, string(readjson.ReadCasesApi(search)))

	} else {
		fmt.Fprintf(w, "You may have tried a different http request than GET or you have not entered a search word")
	}
}
