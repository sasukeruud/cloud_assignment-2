package handlers

import (
	constants "assignment_2/src/app"
	"assignment_2/src/app/structs"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var start time.Time = time.Now()

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		statusHandlerGet(w, r)
	default:
	}
}

func statusHandlerGet(w http.ResponseWriter, r *http.Request) {
	respCases, errCases := http.Get("https://covid19-graphql.now.sh" + "?query=%7B__typename%7D")
	respPolicy, errPolicy := http.Get(constants.CORONA_POLICY_API + "/NOR" + "/2021-01-01")

	if errCases != nil || errPolicy != nil {
		log.Fatal(errCases)
		log.Fatal(errPolicy)
	}

	status := structs.Status{
		CovidCasesApi:  respCases.StatusCode,
		CovidPolicyApi: respPolicy.StatusCode,
		Webhooks:       "test",
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
