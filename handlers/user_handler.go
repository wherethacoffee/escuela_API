package handlers

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

    token, err := middlewares.CreateAccessToken(userFound)
    if err != nil {
	fmt.Println(err)
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Couldn't login",
	})
    }

    cookie := fiber.Cookie{
	Name: "user",
	Value: token,
	Expires: time.Now().Add(time.Hour * 24),
    }

    c.Cookie(&cookie)

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

    token, err := middlewares.CreateAccessToken(userFound)
    if err != nil {
	fmt.Println(err)
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Couldn't login",
	})
    }

    cookie := fiber.Cookie{
	Name: "admin",
	Value: token,
	Expires: time.Now().Add(time.Hour * 24),
	HTTPOnly: true,
    }

    c.Cookie(&cookie)
    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"message": "Login succesfully into admin",
    })
}

func UserProfile(c *fiber.Ctx) error {
   secretKey := os.Getenv("SECRET_KEY")

   if secretKey == "" {
	return fmt.Errorf("Secret key not set in environment variables")
   }

   myKey := []byte(secretKey)

   cookie := c.Cookies("user")

   token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
	return []byte(myKey), nil
   })

   if err != nil {
	return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
	    "error": "Unauthorized",
	})
   }

   claims := token.Claims.(*jwt.RegisteredClaims)

   return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"user token": claims,
   })
}

func AdminProfile(c *fiber.Ctx) error {
   secretKey := os.Getenv("SECRET_KEY")

   if secretKey == "" {
	return fmt.Errorf("Secret key not set in environment variables")
   }

   myKey := []byte(secretKey)

   cookie := c.Cookies("admin")

   token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
	return []byte(myKey), nil
   })

   if err != nil {
	return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
	    "error": "Unauthorized",
	})
   }

   claims := token.Claims.(*jwt.RegisteredClaims)

   return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"admin token": claims,
   })
}

func Logout(c *fiber.Ctx) error {
   cookie := fiber.Cookie{
	Name: "user",
	Value: "",
	Expires: time.Now().Add(-time.Hour),
	HTTPOnly: true,
   }

   c.Cookie(&cookie)

   return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"message": "Logout succesfully",
   })
}

func LogoutAdmin(c *fiber.Ctx) error {
   cookie := fiber.Cookie{
	Name: "admin",
	Value: "",
	Expires: time.Now().Add(-time.Hour),
	HTTPOnly: true,
   }

   c.Cookie(&cookie)

   return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"message": "Logout succesfully",
   })
}
