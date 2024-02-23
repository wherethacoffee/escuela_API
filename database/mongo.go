package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
    usersCollection *mongo.Collection
    recordsCollection *mongo.Collection
    depositsCollection *mongo.Collection
)

//Tests if conncetion to mongo is succesfull 
func Connect(uri, dbName string) error {
    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

    client, err := mongo.Connect(context.TODO(), opts)
    if err != nil {
	panic(err)
    }

    err = client.Ping(context.Background(), nil)
    if err != nil {
	log.Fatal("Error connecting to database:", err)
    }

    fmt.Println("MongoDB connection succesfully")

    usersCollection = client.Database(dbName).Collection("Usuarios")
    recordsCollection = client.Database(dbName).Collection("Registros")
    depositsCollection = client.Database(dbName).Collection("Depositos")

    return nil
}


