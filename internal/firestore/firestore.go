package firestore

import (
	"context"
	"flag"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func CreateClient(ctx context.Context) (*firestore.Client, error) {

	// Sets your Google Cloud Platform project ID
	projectID := ""

	// Override projectID with -project flag
	flag.StringVar(&projectID, "project", projectID, "The Google Cloud Platform project ID.")
	flag.Parse()

	// Create Firestore client
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	} else {
		log.Println("Successfully created Firestore client!")
	}

	return client, nil
}

func AddCollection(client *firestore.Client, ctx context.Context, dataset map[string]interface{}) error {

	// Add dataset into Firestore collection
	_, _, err := client.Collection("convoy").Add(ctx, dataset)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	return nil
}

func ReadCollection(client *firestore.Client, ctx context.Context) (string, error) {

	// Set variables
	var documentID string
	var messageKey = "message"
	var messageValue = "OK!"

	// Read datasets from Firestore collection to find documentID
	iter := client.Collection("convoy").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return "", err
		}
		docData := doc.Data()
		for key, value := range docData {
			if key == messageKey && value == messageValue {
				documentID = doc.Ref.ID
			}
		}
	}

	return documentID, nil
}

func UpdateCollection(client *firestore.Client, ctx context.Context, documentID string) error {

	// Update dataset in Firestore collection
	_, err := client.Collection("convoy").Doc(documentID).Update(ctx, []firestore.Update{
		{
			Path:  "message",
			Value: "OK!",
		},
	})
	if err != nil {
		return err
	}

	return nil
}
