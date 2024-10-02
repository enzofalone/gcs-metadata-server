package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/GoogleCloudPlatform/gcs-metadata-server/internal/repo"
	"github.com/GoogleCloudPlatform/gcs-metadata-server/internal/subscriber"
	"github.com/jessevdk/go-flags"
)

type options struct {
	ProjectId      string `short:"p" long:"project-id" description:"Project ID where subscription resides" required:"true"`
	SubscriptionId string `short:"s" long:"subscription-id" description:"Subscription ID to subscribe to" required:"true"`
	DatabaseUrl    string `short:"d" long:"database-url" description:"Database URL in which to store metadata" required:"true"`
}

func main() {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		os.Exit(1)
	}

	log.Println("Starting subscription service")
	log.Println("Project ID:", opts.ProjectId)
	log.Println("Subscription ID", opts.SubscriptionId)
	log.Println("Database URL:", opts.DatabaseUrl)

	// Connect database
	ctx := context.Background()
	db := repo.NewDatabase(opts.DatabaseUrl, 1)

	if err := db.Connect(ctx); err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}
	defer db.Close()

	if exists, err := db.PingTable(); err != nil || !exists {
		log.Fatalf("Database has not been initialized: %v\n", err)
	}

	// Connect to pub/sub client
	client, err := pubsub.NewClient(context.Background(), opts.SubscriptionId)
	if err != nil {
		log.Fatalf("Error creating pub/sub client: %v\n", err)
	}

	// Start subscription service
	if err := subscriber.Start(client, opts.SubscriptionId, db); err != nil {
		log.Fatalf("Error while listening to subscription: %v\n", err)
	}
}
