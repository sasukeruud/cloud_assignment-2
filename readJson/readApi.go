package readjson

import (
	constants "assignment_2"
	"assignment_2/structs"
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
function to read out the json data from the api.
Takes in one variable to search for a spesific country.
Returns a byte slice of the data*/
func ReadCasesApi(search string) []byte {
	jsonData := map[string]string{
		"query": `
			{
				country(name: "` + search + `"){
					name
					mostRecent{
						date(format: "yyyy-MM-dd")
						confirmed
						deaths
						recovered
						growthRate
					}
				}
			}`,
	}

	//https://www.thepolyglotdeveloper.com/2020/02/interacting-with-a-graphql-api-with-golang/
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

/*
Function to read out data from the policy api.
It takes in to string variabales to define what it will search for.
Returns a slice of policy*/
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
