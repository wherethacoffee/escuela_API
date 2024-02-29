package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/wherethacoffe/escuela_API/database"
	"github.com/wherethacoffe/escuela_API/routes"
)

func main() {
	//Environment variables
	port := os.Getenv("PORT")
	uri := os.Getenv("MONGO_URI")
	db_name := os.Getenv("MONGO_DBNAME")
	
	if port == "" || uri == "" || db_name == "" {
	   fmt.Println("Environment variables not valid") 
	}

	//Fiber app creation
	app := fiber.New()

	//MongoDB connection
	database.Connect(uri, db_name)

	//CORS implementation
	app.Use(cors.New(cors.Config{
	    AllowCredentials: true,
	    AllowOrigins: "http://localhost:5173",
	    AllowMethods: "GET, POST, PUT, DELETE",
	    AllowHeaders: "Origin, Content-Type, Accept",
	}))

	//Routes
	app.Get("/", func(c *fiber.Ctx) error {
	    return c.SendString("Helo, world! 👋")
	})
	routes.UsersRoutes(app)
	routes.RecordsRoutes(app)
	routes.DepositsRoutes(app)

	//Port
	app.Listen(":" + port)
	fmt.Println("Server on port: " + port)
}
