package lib

import (
	"encoding/json"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"cloud.google.com/go/storage"
)
projectID := ""

func dbConnect interface{} {
	ctx := context.Background()

	creds, err := google.FindDefaultCredentials(ctx, storage.ScopeReadOnly)
	if err != nil {
		log.Fatalf(err)
	}
	client, err := storage.NewClient(ctx, option.WithCredentials(creds))
	if err != nil {
		log.Fatalf(err)
	}
	fmt.Println("Buckets: ")
	it := client.Buckets(ctx, projectID)
	for {
		battrs, err := it.Next()
		if err == iterator.Done{
			break
		}
		if err != nil {
			log.Fatalf(err)
		}
		fmt.Println(battrs.Name)
	}
	return client
}
func GetAllUsers() interface{} {
	client := dbConnect()

	iter := client.Collection("users").Documents(ctx)
	for {
		dox, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
}
