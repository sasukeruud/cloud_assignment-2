package handlers

import (
	constants "assignment_2"
	"net/http"
)

/*
Default web handler that gives a short descritption on how to use the api*/
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.Error(w, "To use the api use the following URL's \n"+constants.CASES_PATH+"\n"+
			constants.POLICY_PATH+"\n"+constants.NOTIFICATION_PATH+"\n"+constants.STATUS_PATH, http.StatusAccepted)
	}
}
