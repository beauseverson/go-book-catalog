package testutils

import (
	"context"
	"log"

	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

func SetupTestMongoDB() (*mongodb.MongoDBContainer, string, error) {
	ctx := context.Background()
	mongoContainer, err := mongodb.Run(ctx, "mongo:latest")
	if err != nil {
		return nil, "", err
	}
	uri, err := mongoContainer.ConnectionString(ctx)
	if err != nil {
		return nil, "", err
	}
	log.Println("Test MongoDB URI:", uri)
	return mongoContainer, uri, nil
}
