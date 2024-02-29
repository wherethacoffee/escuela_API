package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wherethacoffe/escuela_API/handlers"
	"github.com/wherethacoffe/escuela_API/middlewares"
)

func RecordsRoutes(app *fiber.App) {
    //Insert user data 
    app.Post("/records/add", middlewares.ValidateToken("user"), handlers.AddRecord)

    //Fetch data from one single user
    app.Get("/records/:_id", middlewares.ValidateToken("user"), handlers.GetRecord)

    //Fetch all records data 
    app.Get("/records", middlewares.ValidateToken("user"), handlers.GetRecords)

    //Update a single user data
    app.Put("/records/update/:_id", middlewares.ValidateToken("user"), handlers.UpdateRecord)

    //Delete a single user data 
    app.Delete("/records/delete/:_id", middlewares.ValidateToken("user"), handlers.DeleteRecord)
}
