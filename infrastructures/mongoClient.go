package infrastructures

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoClient() *mongo.Client {
	client, _ := mongo.Connect(GetConnectionContext(), options.Client().
		ApplyURI("mongodb+srv://joef295:lapices18@cluster0-kgntz.gcp.mongodb.net"))
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connected to MongoDB!")
	return client
}

func GetConnectionContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}
