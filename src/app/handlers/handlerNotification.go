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

}

func postNotification(w http.ResponseWriter, r *http.Request) {
	var webhooks []structs.Webhooks

	webhook := structs.Webhooks{}

	//content, err := json.Marshal()
	AddMessage(w, r)

	err := json.NewDecoder(r.Body).Decode(&webhook)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
	}

	webhooks = append(webhooks, webhook)

	fmt.Println(webhooks)
}

func AddMessage(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./robinruassignment-2-firebase-adminsdk-7fl5y-7ff7b94aac.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	var n structs.Webhooks

	if err != nil {
		log.Fatal("error initializing app:", err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatal(err)
	}
	// very generic way of reading body; should be customized to specific use case
	text, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(text, &n)
	if err != nil {
		http.Error(w, "Reading of payload failed", http.StatusInternalServerError)
		return
	}
	fmt.Println("Received message ", string(text))
	if len(string(text)) == 0 {
		http.Error(w, "Your message appears to be empty. Ensure to terminate URI with /.", http.StatusBadRequest)
	} else {
		// Add element in embedded structure.
		// Note: this structure is defined by the client; but exemplifying a complex one here (including Firestore timestamps).
		id, _, err := client.Collection(collection).Add(ctx,
			map[string]interface{}{
				"url":     n.Url,
				"country": n.Country,
				"calls":   n.Calls,
			})
		//ct++
		if err != nil {
			// Error handling
			http.Error(w, "Error when adding message "+string(text)+", Error: "+err.Error(), http.StatusBadRequest)
			return
		} else {
			fmt.Println("Entry added to collection. Identifier of returned document: " + id.ID)
			// Returns document ID in body
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
