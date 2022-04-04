package firebase

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func InitFirebase() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./robinruassignment-2-firebase-adminsdk-7fl5y-7ff7b94aac.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		fmt.Errorf("error initializing app: %v", err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := client.Close()
		if err != nil {
			log.Fatal("Closing of the firebase client failed. Error:", err)
		}
	}()

	fmt.Println("Firebase conection started")
}
