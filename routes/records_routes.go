package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/wherethacoffe/escuela_API/handlers"
)

func RecordsRoutes(app *fiber.App) {
    //Insert user data 
    app.Post("/records/add", handlers.AddRecord)

    //Fetch data from one single user
    app.Get("/records/:_id", handlers.GetRecord)

    //Fetch all records data 
    app.Get("/records", handlers.GetRecords)

    //Update a single user data
    app.Put("/records/update/:_id", handlers.UpdateRecord)

    //Delete a single user data 
    app.Delete("/records/delete/:_id", handlers.DeleteRecord)
}
