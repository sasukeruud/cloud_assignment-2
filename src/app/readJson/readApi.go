package readjson

import (
	constants "assignment_2/src/app"
	"assignment_2/src/app/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/*
function to read out the json data from the api*/
func ReadCasesApi(search string) []structs.Cases {
	var casesInfo []structs.Cases

	response, err := http.Get(constants.COVID_CASES_API + search)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &casesInfo)

	return casesInfo
}
