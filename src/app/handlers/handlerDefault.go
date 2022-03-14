package handlers

import (
	constants "assignment_2/src/app"
	"fmt"
	"net/http"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		fmt.Println("To use the api use the following URL's \n"+constants.CASES_PATH+"\n",
			constants.POLICY_PATH+"\n"+constants.NOTIFICATION_PATH+"\n"+constants.STATUS_PATH)
	}
}
