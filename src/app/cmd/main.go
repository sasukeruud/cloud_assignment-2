package main

import (
	constants "assignment_2/src/app"
	"assignment_2/src/app/handlers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Running")
	fmt.Print(constants.VERSION)

	http.HandleFunc(constants.DEFAULT_PATH, handlers.DefaultHandler)
	http.HandleFunc(constants.CASES_PATH, handlers.CasesHandler)
	//http.HandleFunc(constants.POLICY_PATH, )
	http.HandleFunc(constants.STATUS_PATH, handlers.StatusHandler)
	//http.HandleFunc(constants.NOTIFICATION_PATH, )
	http.ListenAndServe(":8080", nil)
}
