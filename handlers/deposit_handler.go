package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wherethacoffe/escuela_API/database"
	"github.com/wherethacoffe/escuela_API/middlewares"
	"github.com/wherethacoffe/escuela_API/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddDeposit(c *fiber.Ctx) error {
    var deposit models.Deposit

    if err := c.BodyParser(&deposit); err != nil {
	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
	    "error": "Invalid request payload",
	})
    }

    fecha, err := middlewares.ValidateAndConvertDate(deposit.Fecha)
    if err != nil {
	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
	    "error": "Invalid date format",
	})
    }

    deposit.Fecha = fecha
	
    err = database.InsertDeposit(deposit)
    if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Failed to create deposit",
	})
    }

    return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
	"message": "Deposit added succesfully",
    })
}

func GetDeposit(c *fiber.Ctx) error {
    idString := c.Params("_id")

    id, err := primitive.ObjectIDFromHex(idString)
    if err != nil {
	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
	    "error": "Invalid ID",
	})
    }

    res, err := database.GetDepositByID(id)
    if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Failed to fetch deposit data",
	})
    }

    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"deposit": res,
    })
}

func GetDeposits(c *fiber.Ctx) error {
    res, err := database.GetAllDeposits()
     if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Failed to fetch deposits data",
	})
    }

    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"deposits": res,
    })
}

func UpdateDeposit(c *fiber.Ctx) error {
    idString := c.Params("_id")
    var updatedDeposit models.Deposit

    if err := c.BodyParser(&updatedDeposit); err != nil {
	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
	    "error": "Invalid request payload",
	})
    }

    id, err := primitive.ObjectIDFromHex(idString)
    if err != nil {
	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
	    "error": "Invalid ID",
	})
    }

    fecha, err := middlewares.ValidateAndConvertDate(updatedDeposit.Fecha)
    if err != nil {
	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
	    "error": "Invalid date format",
	})
    }

    updatedDeposit.Fecha = fecha

    err = database.UpdateDepositByID(id, updatedDeposit)
    if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Failed to update deposit",
	})
    }

    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"message": "Deposit updated succesfully",
    })
}

func DeleteDeposit(c *fiber.Ctx) error {
    idString := c.Params("_id")

    id, err := primitive.ObjectIDFromHex(idString)
    if err != nil {
	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
	    "error": "Invalid ID",
	})
    }

    err = database.DeleteDepositByID(id)
    if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Failed to delete deposit",
	})
    }

    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"message": "Deposit deleted succesfully",
    })

}
