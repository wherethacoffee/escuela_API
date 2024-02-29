package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wherethacoffe/escuela_API/handlers"
	"github.com/wherethacoffe/escuela_API/middlewares"
)

func DepositsRoutes(app *fiber.App) {
    //Insert deposit data 
    app.Post("/deposits/add", middlewares.ValidateToken("user"), handlers.AddDeposit)

    //Fetch data from one single deposit
    app.Get("/deposits/:_id", middlewares.ValidateToken("user"), handlers.GetDeposit)

    //Fetch all deposits data 
    app.Get("/deposits", middlewares.ValidateToken("user"), handlers.GetDeposits)

    //Update a single deposit data
    app.Put("/deposits/update/:_id", middlewares.ValidateToken("user"), handlers.UpdateDeposit)

    //Delete a single deposit data 
    app.Delete("/deposits/delete/:_id", middlewares.ValidateToken("user"), handlers.DeleteDeposit)
}
