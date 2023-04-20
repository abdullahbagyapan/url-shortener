package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type DBObject struct {
	ShortURL string
	LongURL  string
}

type Adapter struct {
	db *mongo.Collection
}

func NewAdapter() *Adapter {

	ctx, _ := context.WithCancelCause(context.Background())

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatalf("Error connecting mongodb %v", err)
	}

	collection := client.Database("urlshortener").Collection("urls")

	log.Printf("Succesfully connected to the MongoDB")

	return &Adapter{db: collection}

}

func (da Adapter) SaveURL(longURL, shortURL string) error {

	_, err := da.db.InsertOne(context.Background(), bson.M{"_id": shortURL, "longURL": longURL})

	return err
}

func (da Adapter) GetURL(shortURL string) (string, error) {
	var resp = DBObject{}

	filter := bson.D{{"_id", shortURL}}
	err := da.db.FindOne(context.Background(), filter).Decode(&resp)

	if err != nil {
		return "", err
	}

	return resp.LongURL, nil
}

func (da Adapter) DeleteURL(id string) error {

	return nil

}
