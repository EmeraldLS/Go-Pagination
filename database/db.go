package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err)
	}
	var connectionString = os.Getenv("CONNECTION_STRING")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(connectionString)
	client, _ := mongo.Connect(ctx, clientOptions)
	return client

}

func OpenColletion() *mongo.Collection {
	client := DBInstance()
	collection := client.Database("go_search").Collection("products")
	return collection
}
