package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/wherethacoffe/escuela_API/database"
	"github.com/wherethacoffe/escuela_API/handlers"
)

func main() {

	port := os.Getenv("PORT")
	uri := os.Getenv("MONGO_URI")
	db_name := os.Getenv("MONGO_DBNAME")
	
	if port == "" || uri == "" || db_name == "" {
	   fmt.Println("Environment variables not valid") 
	}

	app := fiber.New()

	database.Connect(uri, db_name)

	app.Use(cors.New())

	//Routes

	//Insert user data 
	app.Post("/users/add", handlers.AddUser)

	//Fetch data from one single user
	app.Get("/users/:username", handlers.GetUser)

	//Fetch all users data 
	app.Get("/users", handlers.GetUsers)

	//Update a single user data
	app.Put("/users/update/:username", handlers.UpdateUser)

	//Delete a single user data 
	app.Delete("/users/delete/:username", handlers.DeleteUser)

	//Login
	app.Post("/login", handlers.Login)

	app.Listen(":" + port)
	fmt.Println("Server on port: " + port)
}
