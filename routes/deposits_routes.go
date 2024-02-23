package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/wherethacoffe/escuela_API/handlers"
)

func DepositsRoutes(app *fiber.App) {
    //Insert deposit data 
    app.Post("/deposits/add", handlers.AddDeposit)

    //Fetch data from one single deposit
    app.Get("/deposits/:_id", handlers.GetDeposit)

    //Fetch all deposits data 
    app.Get("/deposits", handlers.GetDeposits)

    //Update a single deposit data
    app.Put("/deposits/update/:_id", handlers.UpdateDeposit)

    //Delete a single deposit data 
    app.Delete("/deposits/delete/:_id", handlers.DeleteDeposit)
}
