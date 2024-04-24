package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/wherethacoffe/escuela_API/handlers"
    "github.com/wherethacoffe/escuela_API/middlewares"
)

func RecordsRoutes(app *fiber.App) {
    //Insert record data 
    app.Post("/api/records/add", middlewares.ValidateToken("user"), handlers.AddRecord)

    //Fetch data from one single record
    app.Get("/api/records/:_id", middlewares.ValidateToken("user"), handlers.GetRecord)

    //Fetch all records data 
    app.Get("/api/records", middlewares.ValidateToken("user"), handlers.GetRecords)

    //Update a single record data
    app.Put("/api/records/update/:_id", middlewares.ValidateToken("user"), handlers.UpdateRecord)

    //Delete a single record data 
    app.Delete("/api/records/delete/:_id", middlewares.ValidateToken("user"), handlers.DeleteRecord)
}

