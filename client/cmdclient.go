package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

func main() {
	private_key_path := os.Args[1]

	fmt.Println("Hello, World!")

	ctx := context.Background()
	app := initializeAppWithCredential(private_key_path)

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
}
