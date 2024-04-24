package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/wherethacoffe/escuela_API/handlers"
    "github.com/wherethacoffe/escuela_API/middlewares"
)

func UsersRoutes(app *fiber.App) {
    //Insert user data 
    app.Post("/api/users/add", middlewares.ValidateToken("admin"), handlers.AddUser)

    //Fetch data from one single user
    app.Get("/api/users/:username", middlewares.ValidateToken("admin"), handlers.GetUser)

    //Fetch all users data 
    app.Get("/api/users", middlewares.ValidateToken("admin"), handlers.GetUsers)

    //Update a single user data
    app.Put("/api/users/update/:username", middlewares.ValidateToken("admin"), handlers.UpdateUser)

    //Delete a single user data 
    app.Delete("/api/users/delete/:username", middlewares.ValidateToken("admin"), handlers.DeleteUser)

    //Login
    app.Post("/api/login", handlers.Login)

    //Login as admin
    app.Post("/api/loginAdmin", handlers.LoginAdmin)

    //User profile
    app.Get("/api/profile/user", handlers.UserProfile)

    //admin profile
    app.Get("/api/profile/admin", handlers.AdminProfile)

    //Logout
    app.Post("/api/logout", handlers.Logout)

    //Logout as admin
    app.Post("/api/logoutAdmin", handlers.LogoutAdmin)
}

