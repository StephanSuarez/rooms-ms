package conf

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbInstance *mongo.Database
var lock = &sync.Mutex{}

func mongoConnection() {
	mongoURI := "mongodb+srv://stephan2:Atlas123@atlascluster.sexp42w.mongodb.net/"
	opts := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	dbInstance = client.Database("chat-rooms")
}

func GetDBInstance() *mongo.Database {
	if dbInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		mongoConnection()
	}
	return dbInstance
}
