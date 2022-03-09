package main

import (
	constants "assignment_2/src/app"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Running")
	fmt.Print(constants.VERSION)

	//http.HandleFunc(constants.DEFAULT_PATH, )
	//http.HandleFunc(constants.CASES_PATH)
	//http.HandleFunc(constants.POLICY_PATH, )
	//http.HandleFunc(constants.STATUS_PATH, )
	//http.HandleFunc(constants.NOTIFICATION_PATH, )
	http.ListenAndServe(":8080", nil)
}
