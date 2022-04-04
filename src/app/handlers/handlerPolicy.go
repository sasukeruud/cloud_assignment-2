package handlers

import (
	readjson "assignment_2/src/app/readJson"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func PolicyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPolicyRequest(w, r)
	default:
		fmt.Println("test")
	}
}

func getPolicyRequest(w http.ResponseWriter, r *http.Request) {
	search := strings.SplitAfter(r.URL.Path, "/")

	if search[len(search)-1] != "policy" {
		w.Header().Set("content-type", "application/json")

		encoder := json.NewEncoder(w)

		err := encoder.Encode(readjson.ReadPolicyApi(search[4], search[5]))
		if err != nil {
			http.Error(w, "Error during encoding", http.StatusInternalServerError)
		}

	} else {
		fmt.Fprintf(w, "You may have tried a different http request than GET or you have not entered a search word")
	}
}
