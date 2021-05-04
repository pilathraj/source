package helper

import (
	"context"
	"sync"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
/* Used to create a singleton object of MongoDB client.
Initialized and exposed through  GetMongoClient().*/
var clientInstance *mongo.Client
//Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error
//Used to execute client creation procedure only once.
var mongoOnce sync.Once
//I have used below constants just to hold required database config's.
const (
	CONNECTIONSTRING = "mongodb+srv://demo:demo123@cluster0.5bxx7.mongodb.net/talentpro?retryWrites=true&w=majority"
	DB               = "talentpro"
	COLLECTION           = "col_users"
)
//GetMongoClient - Return mongodb connection to work with
func GetMongoClient() (*mongo.Client, error) {
	//Perform connection creation operation only once.
	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}

func GetCollection() *mongo.Collection {
	client, err := GetMongoClient()
	if err != nil {
		log.Println(err)		
	}
	return client.Database(DB).Collection(COLLECTION)
}