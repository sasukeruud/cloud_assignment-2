package handlers

import (
	"assignment_2/src/app/structs"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

// Collection name in Firestore
const collection = "webhooks"
const coll = "country_calls"

func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	ctx = context.Background()
	opt := option.WithCredentialsFile("./robinruassignment-2-firebase-adminsdk-7fl5y-7ff7b94aac.json")
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

func getNotification(w http.ResponseWriter, r *http.Request) {
	var webhooks []structs.Webhooks
	var o structs.Webhooks
	elem := strings.Split(r.URL.Path, "/")

	if len(elem) > 3 {
		search := elem[4]
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

func deleteNotification(w http.ResponseWriter, r *http.Request) {
	elem := strings.Split(r.URL.Path, "/")

	if len(elem) >= 4 {
		delete := elem[4]
		res, err := client.Collection(collection).Doc(delete).Delete(ctx)
		if err != nil {
			fmt.Fprint(w, "error when trying to delete")
			return
		}

		fmt.Fprintf(w, "webhook was deleted at: "+res.UpdateTime.GoString())

	} else {

	}
}

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
				"calls":   o.Calls,
			})
		if err != nil {
			http.Error(w, "Error when adding message "+string(info)+", Error: "+err.Error(), http.StatusBadRequest)
			return
		} else {
			http.Error(w, id.ID, http.StatusCreated)
			return
		}
	}
}

func GetWebhooks(w http.ResponseWriter, r *http.Request) []structs.Webhooks {
	var webhooks []structs.Webhooks
	var o structs.Webhooks

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

	return webhooks
}

func WebhookCall(w http.ResponseWriter, r *http.Request, search string) {
	webhooks = GetWebhooks(w, r)
	country_calls := structs.Country_calls{}

	res := client.Collection(coll).Doc(search)

	doc, err := res.Get(ctx)
	if err != nil {
		http.Error(w, "Error extracting body of returned document of message "+search, http.StatusInternalServerError)
		return
	}

	jsonString, err := json.Marshal(doc.Data())
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(jsonString, &country_calls)

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
				go CallUrl(e.Url, "POST", jsonString)
				country_calls.Called = country_calls.Called + 1
			} else {
				country_calls.Called = country_calls.Called + 1
			}
		}

	}

	client.Collection(coll).Doc(search).Set(ctx, map[string]interface{}{
		"called": country_calls.Called,
	})
}

func CallUrl(url string, method string, content []byte) {
	//fmt.Println("Attempting invocation of url " + url + " with content '" + content + "'.")
	//res, err := http.Post(url, "text/plain", bytes.NewReader([]byte(content)))
	req, err := http.NewRequest(method, url, bytes.NewReader(content))
	if err != nil {
		log.Printf("%v", "Error during request creation. Error:", err)
		return
	}

	// Perform invocation
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error in HTTP request. Error:", err)
		return
	}

	// Read the response
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Something is wrong with invocation response. Error:", err)
		return
	}

	fmt.Println("Webhook invoked. Received status code " + strconv.Itoa(res.StatusCode) +
		" and body: " + string(response))
}
