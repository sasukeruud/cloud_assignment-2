package handlers

import (
	constants "assignment_2"
	"assignment_2/structs"
	"encoding/json"
	"net/http"
	"time"
)

//Variable to record how long the application have been running
var start time.Time = time.Now()

/*
Function with a switch-case to handle what function to be executed by which http request*/
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		statusHandlerGet(w, r)
	default:
		http.Error(w, "Only http GET request is implemented on this endpoint", http.StatusOK)
	}
}

/*
Function to create an object of stucts.status that will be encoded to json
and give information about the application to the user.*/
func statusHandlerGet(w http.ResponseWriter, r *http.Request) {
	respCases, errCases := http.Get("https://covid19-graphql.now.sh" + "?query=%7B__typename%7D")
	respPolicy, errPolicy := http.Get(constants.CORONA_POLICY_API + "/NOR" + "/2021-01-01")

	if errCases != nil || errPolicy != nil {
		http.Error(w, errCases.Error(), http.StatusNotFound)
		http.Error(w, errPolicy.Error(), http.StatusNotFound)
	}

	status := structs.Status{
		CovidCasesApi:  respCases.StatusCode,
		CovidPolicyApi: respPolicy.StatusCode,
		Webhooks:       len(GetWebhooks(w, r)),
		Version:        constants.VERSION,
		Uptime:         time.Duration.Seconds(time.Since(start)),
	}

	w.Header().Set("content-type", "application/json")

	encoder := json.NewEncoder(w)

	err := encoder.Encode(status)

	if err != nil {
		http.Error(w, "error during encoding", http.StatusInternalServerError)
	}
}
