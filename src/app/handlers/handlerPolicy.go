package handlers

import (
	readjson "assignment_2/src/app/readJson"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
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
	search := path.Base(r.URL.Path)

	if search != "cases" {
		w.Header().Set("contet-type", "application/json")

		encoder := json.NewEncoder(w)

		err := encoder.Encode(readjson.ReadPolicyApi(search))
		if err != nil {
			http.Error(w, "Error during encoding", http.StatusInternalServerError)
		}

	} else {
		fmt.Println("You may have tried a different http request than GET or you have not entered a search word")
	}
}
