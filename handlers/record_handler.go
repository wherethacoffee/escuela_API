package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wherethacoffe/escuela_API/database"
	"github.com/wherethacoffe/escuela_API/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddRecord(c *fiber.Ctx) error {
    var record models.Record

    if err := c.BodyParser(&record); err != nil {
	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
	    "error": "Invalid request payload",
	})
    }

    record.Total = record.Colegiatura + record.Extra

    err := database.InsertRecord(record)
    if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Failed to create record",
	})
    }

    return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
	"message": "Record added succesfully",
    })
}

func GetRecord(c *fiber.Ctx) error {
    idString := c.Params("_id")

    id, err := primitive.ObjectIDFromHex(idString)
    if err != nil {
	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
	    "error": "Invalid ID",
	})
    }

    res, err := database.GetRecordByID(id)
    if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Failed to fetch record data",
	})
    }

    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"record": res,
    })
}

func GetRecords(c *fiber.Ctx) error {
    res, err := database.GetAllRecords()
     if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Failed to fetch records data",
	})
    }

    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"records": res,
    })
}

func UpdateRecord(c *fiber.Ctx) error {
    idString := c.Params("_id")
    var updatedRecord models.Record

    if err := c.BodyParser(&updatedRecord); err != nil {
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


    updatedRecord.Total = updatedRecord.Colegiatura + updatedRecord.Extra

    err = database.UpdateRecordByID(id, updatedRecord)
    if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Failed to update record",
	})
    }

    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"message": "Record updated succesfully",
    })
}

func DeleteRecord(c *fiber.Ctx) error {
    idString := c.Params("_id")

    id, err := primitive.ObjectIDFromHex(idString)
    if err != nil {
	return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
	    "error": "Invalid ID",
	})
    }

    err = database.DeleteRecordByID(id)
    if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
	    "error": "Failed to delete record",
	})
    }

    return c.Status(fiber.StatusOK).JSON(&fiber.Map{
	"message": "Record deleted succesfully",
    })

}
