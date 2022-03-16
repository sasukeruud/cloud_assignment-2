package handlers

import (
	readjson "assignment_2/src/app/readJson"
	"encoding/json"
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
		fmt.Println("The only http method implemented is the GET request")
	}
}

func casesGetRequest(w http.ResponseWriter, r *http.Request) {
	search := path.Base(r.URL.Path)

	if search != "cases" {
		w.Header().Set("contet-type", "application/json")

		encoder := json.NewEncoder(w)

		err := encoder.Encode(readjson.ReadCasesApi(search))
		if err != nil {
			http.Error(w, "Error during encoding", http.StatusInternalServerError)
		}

	} else {
		fmt.Println("You may have tried a different http request than GET or you have not entered a search word")
	}
}
