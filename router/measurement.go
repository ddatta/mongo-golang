package router



import (
	"fmt"
	"github.com/ddatta/mongo-golang/common"
	"github.com/ddatta/mongo-golang/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createDTO struct {
	VehicleId  int `json:"vehicleId" bson:"vehicleId"`
	// Author string `json:"author" bson:"author"`
	// Year   string `json:"year" bson:"year"`
}

func AddMeasurementGroup(app *fiber.App) {
	measurementGroup := app.Group("/api")

	measurementGroup.Get("/measurements", getMeasurements)
	measurementGroup.Get("/measurement/:id", getMeasurement)
	measurementGroup.Post("/measurement", createMeasurement)
//	measurementGroup.Put("/:id", updateMeasurement)
//	measurementGroup.Delete("/:id", deleteMeasurement)
}

func getMeasurements(c *fiber.Ctx) error {
	fmt.Println("Here ... .")
	coll := common.GetDBCollection("measurement")

	// find all 
	measurements := make([]models.Measurement, 0)
	cursor, err := coll.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// iterate over the cursor
	for cursor.Next(c.Context()) {
		measurement := models.Measurement{}
		err := cursor.Decode(&measurement)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		measurements = append(measurements, measurement)
	}

	return c.Status(200).JSON(fiber.Map{"data": measurements})
}

func getMeasurement(c *fiber.Ctx) error {
	coll := common.GetDBCollection("measurement")

	// find the measurement
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	measurement := models.Measurement{}

	err = coll.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&measurement)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{"data": measurement})
}



func createMeasurement(c *fiber.Ctx) error {
	// validate the body
	b := new(createDTO)
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	// create the measurement
	coll := common.GetDBCollection("measurement")
	result, err := coll.InsertOne(c.Context(), b)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create measurement",
			"message": err.Error(),
		})
	}

	// return the measurement
	return c.Status(201).JSON(fiber.Map{
		"result": result,
	})
}

type updateDTO struct {
	VehicleId  string `json:"vehicleId,omitempty" bson:"vehicleId,omitempty"`
	// Author string `json:"author,omitempty" bson:"author,omitempty"`
	// Year   string `json:"year,omitempty" bson:"year,omitempty"`
}

// func updateMeasurement(c *fiber.Ctx) error {
// 	// validate the body
// 	b := new(updateDTO)
// 	if err := c.BodyParser(b); err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "Invalid body",
// 		})
// 	}

// 	// get the id
// 	id := c.Params("id")
// 	if id == "" {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "id is required",
// 		})
// 	}
// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "invalid id",
// 		})
// 	}

// 	// update the measurement
// 	coll := common.GetDBCollection("books")
// 	result, err := coll.UpdateOne(c.Context(), bson.M{"_id": objectId}, bson.M{"$set": b})
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"error":   "Failed to update book",
// 			"message": err.Error(),
// 		})
// 	}

// 	// return the book
// 	return c.Status(200).JSON(fiber.Map{
// 		"result": result,
// 	})
// }

// func deleteBook(c *fiber.Ctx) error {
// 	// get the id
// 	id := c.Params("id")
// 	if id == "" {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "id is required",
// 		})
// 	}
// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "invalid id",
// 		})
// 	}

// 	// delete the book
// 	coll := common.GetDBCollection("books")
// 	result, err := coll.DeleteOne(c.Context(), bson.M{"_id": objectId})
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"error":   "Failed to delete book",
// 			"message": err.Error(),
// 		})
// 	}

// 	return c.Status(200).JSON(fiber.Map{
// 		"result": result,
// 	})
// }