package main

import (
		"context"
		"encoding/json"
		"os" 
	
	"github.com/joho/godotenv"
	"o.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mogo"
	"go.mongodb.or/mongo-driver/mongo/options"
)


func Clinet () *mongo.Client{
	if err := godotenv.Load(); err != nil {
		log.Println("N .env file found")
	}
	ui := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental ariable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	clent, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if rr != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	//col := client.Database("wrapme-local").Collection("contacts")
	return client
}
