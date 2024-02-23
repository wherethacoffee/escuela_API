package handlers

import (

	"github.com/gofiber/fiber/v2"
	"github.com/wherethacoffe/escuela_API/database"
	"github.com/wherethacoffe/escuela_API/middlewares"
	"github.com/wherethacoffe/escuela_API/models"
)

func AddUser(c *fiber.Ctx) error {
    var user models.User

    if err := c.BodyParser(&user); err != nil {
	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
	    "error": "Invalid request payload",
	})
    }

    hash, err := middlewares.HashPassword(user.Password)
    if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Failed to create user",
	})
    }

    user.Password = hash

    err = database.InsertUser(user)
    if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Failed to create user",
	})
    }

    return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
	"message": "User added succesfully",
    })
}

func GetUser(c *fiber.Ctx) error {
    username := c.Params("username")

    res, err := database.GetUserByUsername(username)
    if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Failed to fetch user data",
	})
    }

    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"user": res,
    })
}

func GetUsers(c *fiber.Ctx) error {
    res, err := database.GetAllUsers()
     if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Failed to fetch users data",
	})
    }

    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"users": res,
    })
}

func UpdateUser(c *fiber.Ctx) error {
    username := c.Params("username")
    var updatedUser models.User

    if err := c.BodyParser(&updatedUser); err != nil {
	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
	    "error": "Invalid request payload",
	})
    }

    hash, err := middlewares.HashPassword(updatedUser.Password)
    if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Failed to update user",
	})
    }

    updatedUser.Password = hash

    err = database.UpdateUserByUsername(username, updatedUser)
    if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Failed to update user",
	})
    }

    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"message": "User updated succesfully",
    })
}

func DeleteUser(c *fiber.Ctx) error {
    username := c.Params("username")

    err := database.DeleteUserByUsername(username)
    if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Failed to delete user",
	})
    }

    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"message": "User deleted succesfully",
    })

}

func Login(c *fiber.Ctx) error {
    var user models.User

    if err := c.BodyParser(&user); err != nil {
	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
	    "error": "Invalid request payload",
	})
    }

    userFound, err := database.GetUserByUsername(user.Username)
    if err != nil {
	return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
	    "error": "Failed to fetch user",
	})
    }
    isMatch, err := middlewares.CheckPassword(user.Password, userFound.Password)
    if !isMatch || err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Wrong password",
	})
    }
    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"message": "Login succesfully",
    })
}

func LoginAdmin(c *fiber.Ctx) error {
    var user models.User

    if err := c.BodyParser(&user); err != nil {
	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
	    "error": "Invalid request payload",
	})
    }

    userFound, err := database.GetUserByUsername(user.Username)
    if err != nil {
	return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
	    "error": "Failed to fetch user",
	})
    }
    isMatch, err := middlewares.CheckPassword(user.Password, userFound.Password)
    if !isMatch || err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Wrong password",
	})
    }
    isAdmin := userFound.IsAdmin
    if !isAdmin {
	return c.Status(fiber.StatusForbidden).JSON(&fiber.Map{
	    "error": "You have not admin credentials",
	})
    }
    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"message": "Login succesfully into admin",
    })
}
