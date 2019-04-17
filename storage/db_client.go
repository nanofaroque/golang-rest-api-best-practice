package storage

import (
	"context"
	"fmt"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)
func GetClient() *mongo.Client {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err!= nil{
		return nil
	}
	fmt.Print(client)
	return client
}

func GetCollection() *mongo.Collection {
	return GetClient().Database("testing").Collection("numbers")
}
