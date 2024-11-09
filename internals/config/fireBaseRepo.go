package config

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"

	"google.golang.org/api/option"
)

var (
	ctx    context.Context
	client *firestore.Client
)

func fireBaseConfig() {
	wd, _ := os.Getwd()
	envPath := filepath.Join(wd, "../../.env")
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("unable to load env file: %v", err)
	}
	ctx = context.Background()

	opt := option.WithCredentialsFile(os.Getenv("PATHFILE"))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Firebase initialization failed: %v", err)
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Firestore initialization failed: %v", err)
	}
}

func CloseFirestoreClient() {
	if err := client.Close(); err != nil {
		log.Fatalf("Failed to close Firestore client: %v", err)
	}
}
