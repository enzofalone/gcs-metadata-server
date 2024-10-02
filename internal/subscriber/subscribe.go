package subscriber

import (
	"cloud.google.com/go/pubsub"
	"github.com/GoogleCloudPlatform/gcs-metadata-server/internal/repo"
)

func Start(client *pubsub.Client, subscriptionId string, db *repo.Database) error {
	return nil
}
