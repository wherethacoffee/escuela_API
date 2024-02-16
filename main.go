package main

import (
	"context"
	"fmt"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	port := os.Getenv("PORT")
	uri := os.Getenv("MONGO_URI")
	db_name := os.Getenv("MONGO_DBNAME")
	
	if port == "" || uri == "" || db_name == "" {
	   fmt.Println("Environment variables not valid") 
	}

	app := fiber.New()


	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
	   panic(err)
	}

	coll := client.Database(db_name).Collection("admins")

	//user := coll.FindOne(context.TODO(), bson.D{{Key: "username", Value: "adminMaster"}})
	coll.InsertOne(context.TODO(), bson.D{
	   {Key: "username",
	    Value: "adminTestGo3",},
	   {Key: "pwd",
	    Value: "gotest123_3",},
    	})

	

	app.Use(cors.New())



	app.Get("/users", func(c *fiber.Ctx) error {
	   return c.JSON(&fiber.Map{
	      "data": "backend users",
	  })
	})

	app.Post("/users", func(c *fiber.Ctx) error {
	   return c.JSON(&fiber.Map{
	      "data": "adding user",
	   })
	})

	app.Listen(":" + port)
	fmt.Println("Server on port: " + port)
}
