package readjson

import (
	constants "assignment_2/src/app"
	"assignment_2/src/app/structs"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

/*
function to read out the json data from the api*/
func ReadCasesApi(search string) []byte {
	jsonData := map[string]string{
		"query": `
			{
				country(name: "` + search + `"){
					name
					mostRecent{
						date(format: "yyyy-MM-dd")
						confirmed
					}
				}
			}`,
	}

	jsonValue, _ := json.Marshal(jsonData)
	request, err := http.NewRequest("POST", constants.COVID_CASES_API, bytes.NewBuffer(jsonValue))
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)

	return data
}

func ReadPolicyApi(country, date string) []structs.Policy {
	var policyInfo []structs.Policy
	var policy structs.Policy

	response, err := http.Get(constants.CORONA_POLICY_API + country + date)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &policy)

	policyInfo = append(policyInfo, policy)

	return policyInfo
}
