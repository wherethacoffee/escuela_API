package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/wherethacoffe/escuela_API/handlers"
)

func UsersRoutes(app *fiber.App) {
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

    //Login como admin
    app.Post("/loginAdmin", handlers.LoginAdmin)
}
