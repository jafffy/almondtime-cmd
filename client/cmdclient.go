package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func initializeAppWithCredential(private_key_path string) *firebase.App {
	opt := option.WithCredentialsFile(private_key_path)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app
}

func initializeFirestoreClient(ctx context.Context, app *firebase.App) *firestore.Client {
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client
}

func addActivityLog(ctx context.Context, client *firestore.Client) {
	_, _, err := client.Collection("logs").Add(ctx, map[string]interface{}{
		"date": time.Now(),
	})
	if err != nil {
		log.Fatalf("Failed adding log: %v", err)
	}
}

func main() {
	private_key_path := os.Args[1]

	fmt.Println("Hello, World!")

	ctx := context.Background()
	app := initializeAppWithCredential(private_key_path)
	firestore_client := initializeFirestoreClient(ctx, app)

	addActivityLog(ctx, firestore_client)

	defer firestore_client.Close()
}
