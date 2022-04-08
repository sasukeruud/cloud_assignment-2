package handlers

import (
	"assignment_2/structs"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// Firebase context and client used by Firestore functions throughout the program.
var ctx context.Context
var client *firestore.Client
var webhooks = []structs.Webhooks{}
var mock = false

// Collection name in Firestore
const collection = "webhooks"
const coll = "country_calls"

/*
Handler to check what type of http request that get sent to the application.
Will execute correct code based on what the http request is*/
func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	ctx = context.Background()
	opt := option.WithCredentialsFile("./auth.json")
	app, err := firebase.NewApp(ctx, nil, opt)

	if err != nil {
		log.Fatal("error initializing app:", err)
	}

	client, err = app.Firestore(ctx)

	if err != nil {
		log.Fatal(err)
	}

	switch r.Method {
	case http.MethodPost:
		postNotification(w, r)
	case http.MethodDelete:
		deleteNotification(w, r)
	case http.MethodGet:
		getNotification(w, r)
	default:
	}

	defer func() {
		err := client.Close()
		if err != nil {
			log.Fatal("Closing of the firebase client failed. Error:", err)
		}
	}()
}

/*
Function that will give a response of either one webhook data or a list of webhook data.*/
func getNotification(w http.ResponseWriter, r *http.Request) {
	var webhooks []structs.Webhooks
	var o structs.Webhooks
	elem := strings.Split(r.URL.Path, "/")

	if len(elem) > 3 {
		search := elem[4]
		//checks that there is data in the fourth element of the slice
		if len(search) != 0 {
			res := client.Collection(collection).Doc(search)

			doc, err := res.Get(ctx)
			if err != nil {
				http.Error(w, "Error extracting body of returned document of message "+search, http.StatusInternalServerError)
				return
			}

			jsonString, err := json.Marshal(doc.Data())
			if err != nil {
				log.Fatal(err)
			}

			json.Unmarshal(jsonString, &o)
			webhooks = append(webhooks, o)
			err1 := json.NewEncoder(w).Encode(webhooks)
			if err1 != nil {
				http.Error(w, "Error during encoding", http.StatusInternalServerError)
			}
		} else {
			iter := client.Collection(collection).Documents(ctx)

			for {
				doc, err := iter.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					log.Fatalf("Failed to iterate: %v", err)
				}

				jsonString, err := json.Marshal(doc.Data())
				if err != nil {
					log.Fatal(err)
				}
				json.Unmarshal(jsonString, &o)
				o.WebhookID = doc.Ref.ID
				webhooks = append(webhooks, o)
			}
			err := json.NewEncoder(w).Encode(webhooks)
			if err != nil {
				http.Error(w, "Error during encoding", http.StatusInternalServerError)
			}
		}

	} else {
		fmt.Fprint(w, "something wrong have happened", http.StatusBadGateway)
	}
}

/*
Function to delete a webhook from the firebase databse.
Uses the webhook ID to check which webhook it will delete*/
func deleteNotification(w http.ResponseWriter, r *http.Request) {
	elem := strings.Split(r.URL.Path, "/")

	if len(elem) >= 4 {
		delete := elem[4]
		res, err := client.Collection(collection).Doc(delete).Delete(ctx)
		if err != nil {
			http.Error(w, "error when trying to delete", http.StatusAccepted)
		}

		fmt.Fprintf(w, "webhook was deleted at: "+res.UpdateTime.GoString())

	} else {

	}
}

/*
Function to create a new set of data in the firebase database.
Takes a json body to define what the data is going to be*/
func postNotification(w http.ResponseWriter, r *http.Request) {
	var o structs.Webhooks
	info, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading firebase database", http.StatusInternalServerError)
		return
	}
	if len(string(info)) == 0 {
		http.Error(w, "You have not entered any inforamation", http.StatusBadRequest)
	} else {
		json.Unmarshal(info, &o)
		id, _, err := client.Collection(collection).Add(ctx,
			map[string]interface{}{
				"url":     o.Url,
				"country": o.Country,
				"calls":   o.Calls})

		if err != nil {
			http.Error(w, "Error when adding message "+string(info)+", Error: "+err.Error(), http.StatusBadRequest)
			return
		} else {
			http.Error(w, id.ID, http.StatusCreated)
			return
		}
	}
}

/*
Function to get the data from the firebase database.
returns a slice of webhooks.*/
func GetWebhooks(w http.ResponseWriter, r *http.Request) []structs.Webhooks {
	if !mock {
		var webhooks []structs.Webhooks
		var o structs.Webhooks
		ctx = context.Background()
		opt := option.WithCredentialsFile("./auth.json")
		app, err := firebase.NewApp(ctx, nil, opt)

		if err != nil {
			log.Fatal("error initializing app:", err)
		}

		client, err = app.Firestore(ctx)

		if err != nil {
			log.Fatal(err)
		}

		iter := client.Collection(collection).Documents(ctx)

		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalf("Failed to iterate: %v", err)
			}

			jsonString, err := json.Marshal(doc.Data())
			if err != nil {
				log.Fatal(err)
			}
			json.Unmarshal(jsonString, &o)
			o.WebhookID = doc.Ref.ID
			webhooks = append(webhooks, o)
		}
		defer func() {
			err := client.Close()
			if err != nil {
				log.Fatal("Closing of the firebase client failed. Error:", err)
			}
		}()

		return webhooks
	} else {
		return FirebaseMock()
	}
}

/*
Function to give out a POST request when a country it called*/
func WebhookCall(w http.ResponseWriter, r *http.Request, search string) {
	if !mock {
		webhooks = GetWebhooks(w, r)
		country_calls := structs.Country_calls{}
		ctx = context.Background()
		opt := option.WithCredentialsFile("./auth.json")
		app, err := firebase.NewApp(ctx, nil, opt)

		if err != nil {
			log.Fatal("error initializing app:", err)
		}

		client, err = app.Firestore(ctx)

		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(getCountry(search), &country_calls)

		str, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal("Error during decoding message content. Error: " + string(str))
		}

		for _, e := range webhooks {
			jsonString, err := json.Marshal(e)
			if err != nil {
				http.Error(w, "something went wrong", http.StatusBadGateway)
			}
			if e.Country == search {
				if country_calls.Called >= e.Calls {
					go callUrl(e.Url, "POST", jsonString)
					country_calls.Called = country_calls.Called + 1
				} else {
					country_calls.Called = country_calls.Called + 1
				}
			}

		}

		client.Collection(coll).Doc(country_calls.Country_id).Set(ctx, map[string]interface{}{
			"country": search,
			"called":  country_calls.Called,
		})
		defer func() {
			err := client.Close()
			if err != nil {
				log.Fatal("Closing of the firebase client failed. Error:", err)
			}
		}()
	} else {
		return
	}
}

func callUrl(url string, method string, content []byte) {
	req, err := http.NewRequest(method, url, bytes.NewReader(content))
	if err != nil {
		log.Printf("%v Error during request creation. Error:", err)
		return
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error in HTTP request. Error:", err)
		return
	}

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Something is wrong with invocation response. Error:", err)
		return
	}

	fmt.Println("Webhook invoked. Received status code " + strconv.Itoa(res.StatusCode) +
		" and body: " + string(response))
}

/*
Function to get information about how many times a country have been called on by the application.
Data is store on firebase.*/
func getCountry(search string) []byte {
	var country structs.Country_calls

	iter := client.Collection(coll).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			id, _, err := client.Collection(collection).Add(ctx,
				map[string]interface{}{
					"country": search,
					"calls":   0})

			if err != nil {
				break
			} else {
				country = structs.Country_calls{id.ID, search, 0}
			}
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		jsonString, err := json.Marshal(doc.Data())
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(jsonString, &country)
		country.Country_id = doc.Ref.ID
		if country.Country == search {
			break
		}
	}

	jsonString, err := json.Marshal(country)
	if err != nil {

	}

	return jsonString
}

/*
Calls a json file to mock some parts of the database*/
func FirebaseMock() []structs.Webhooks {
	var webhook structs.Webhooks
	var webhooks []structs.Webhooks
	jsonFile, err := os.Open("res/firebaseMock.json")
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(byteValue, &webhook)
	webhooks = append(webhooks, webhook)

	return webhooks
}
