package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wherethacoffe/escuela_API/handlers"
	"github.com/wherethacoffe/escuela_API/middlewares"
)

func UsersRoutes(app *fiber.App) {
    //Insert user data 
    app.Post("/users/add", middlewares.ValidateToken("admin"), handlers.AddUser)

    //Fetch data from one single user
    app.Get("/users/:username", middlewares.ValidateToken("admin"), handlers.GetUser)

    //Fetch all users data 
    app.Get("/users", middlewares.ValidateToken("admin"), handlers.GetUsers)

    //Update a single user data
    app.Put("/users/update/:username", middlewares.ValidateToken("admin"), handlers.UpdateUser)

    //Delete a single user data 
    app.Delete("/users/delete/:username", middlewares.ValidateToken("admin"), handlers.DeleteUser)

    //Login
    app.Post("/login", handlers.Login)

    //Login as admin
    app.Post("/loginAdmin", handlers.LoginAdmin)

    //User profile
    app.Get("/profile/user", handlers.UserProfile)

    //admin profile
    app.Get("/profile/admin", handlers.AdminProfile)

    //Logout
    app.Post("/logout", handlers.Logout)

    //Logout as admin
    app.Post("/logoutAdmin", handlers.LogoutAdmin)
}
