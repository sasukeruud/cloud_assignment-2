package handlers

import (
	readjson "assignment_2/readJson"
	"encoding/json"
	"net/http"
	"strings"
)

/*
Function with switch statement to ensure that the right function get executed with the
correct http request.*/
func PolicyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPolicyRequest(w, r)
	default:
		http.Error(w, "Only http GET request is implemented", http.StatusOK)
	}
}

/*
Function to encode the response from the policy api*/
func getPolicyRequest(w http.ResponseWriter, r *http.Request) {
	search := strings.SplitAfter(r.URL.Path, "/")

	if search[len(search)-1] != "policy" && len(search) == 6 {
		w.Header().Set("content-type", "application/json")

		encoder := json.NewEncoder(w)

		err := encoder.Encode(readjson.ReadPolicyApi(search[4], search[5]))
		if err != nil {
			http.Error(w, "Error during encoding", http.StatusInternalServerError)
		}

	} else {
		http.Error(w, "You may have tried a different http request than GET or you have not entered a search word", http.StatusNonAuthoritativeInfo)
	}
}
