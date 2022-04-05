package handlers

import (
	"assignment_2/src/app/structs"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// Firebase context and client used by Firestore functions throughout the program.
var ctx context.Context
var client *firestore.Client

// Collection name in Firestore
const collection = "webhooks"

func NotificationHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		postNotification(w, r)
	case http.MethodDelete:
		deleteNotification(w, r)
	case http.MethodGet:
		getNotification(w, r)
	default:
	}
}

func getNotification(w http.ResponseWriter, r *http.Request) {
	var webhooks []structs.Webhooks
	var o structs.Webhooks

	ctx := context.Background()
	opt := option.WithCredentialsFile("./robinruassignment-2-firebase-adminsdk-7fl5y-7ff7b94aac.json")
	app, err := firebase.NewApp(ctx, nil, opt)

	if err != nil {
		log.Fatal("error initializing app:", err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatal(err)
	}
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

	defer func() {
		err := client.Close()
		if err != nil {
			log.Fatal("Closing of the firebase client failed. Error:", err)
		}
	}()
}

func deleteNotification(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	opt := option.WithCredentialsFile("./robinruassignment-2-firebase-adminsdk-7fl5y-7ff7b94aac.json")
	app, err := firebase.NewApp(ctx, nil, opt)

	if err != nil {
		log.Fatal("error initializing app:", err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatal(err)
	}
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

	defer func() {
		err := client.Close()
		if err != nil {
			log.Fatal("Closing of the firebase client failed. Error:", err)
		}
	}()
}

func postNotification(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./robinruassignment-2-firebase-adminsdk-7fl5y-7ff7b94aac.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	var o structs.Webhooks

	if err != nil {
		log.Fatal("error initializing app:", err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatal(err)
	}

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

	defer func() {
		err := client.Close()
		if err != nil {
			log.Fatal("Closing of the firebase client failed. Error:", err)
		}
	}()
}

func GetWebhookNumber(w http.ResponseWriter, r *http.Request) []structs.Webhooks {
	var webhooks []structs.Webhooks
	var o structs.Webhooks

	ctx := context.Background()
	opt := option.WithCredentialsFile("./robinruassignment-2-firebase-adminsdk-7fl5y-7ff7b94aac.json")
	app, err := firebase.NewApp(ctx, nil, opt)

	if err != nil {
		log.Fatal("error initializing app:", err)
	}

	client, err := app.Firestore(ctx)

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
}
