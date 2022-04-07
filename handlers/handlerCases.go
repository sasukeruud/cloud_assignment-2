package handlers

import (
	readjson "assignment_2/readJson"
	"net/http"
	"path"
)

/*
Function to handle different types of https requests. switch case that handle
 GET request spesificly all other are handled under default.*/
func CasesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		casesGetRequest(w, r)
	default:
		http.Error(w, "The only http method implemented is the GET request", http.StatusOK)
	}
}

/*
Function to give the user a response from a GET request*/
func casesGetRequest(w http.ResponseWriter, r *http.Request) {
	search := path.Base(r.URL.Path)
	if search != "cases" {

		w.Header().Set("content-type", "application/json")

		http.Error(w, string(readjson.ReadCasesApi(search)), http.StatusOK)
		WebhookCall(w, r, search)

	} else {
		http.Error(w, "You may have tried a different http request than GET or you have not entered a search word", http.StatusOK)
	}
}
