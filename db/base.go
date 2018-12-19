package db

import (
	"context"
	"github.com/Echo-Project-Structure/config"
	mongo "github.com/mongodb/mongo-go-driver/mongo"
	"time"
)

// GetMongoDbInstance connect and get connection instance
func GetMongoDbInstance() *mongo.Client {
	client, err := mongo.NewClient(config.MongoURI)
	if err != nil{
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil{
		panic(err)
	}
	return client
}
