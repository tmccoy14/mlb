package firestore

import (
	"context"
	"flag"
	"log"

	"cloud.google.com/go/firestore"
)

func createClient(ctx context.Context) *firestore.Client {

	// Sets your Google Cloud Platform project ID
	projectID := ""

	// Override projectID with -project flag
	flag.StringVar(&projectID, "project", projectID, "The Google Cloud Platform project ID.")
	flag.Parse()

	// Create Firestore client
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	} else {
		log.Println("Successfully created client!")
	}

	return client
}
