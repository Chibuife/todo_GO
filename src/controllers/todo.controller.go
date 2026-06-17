package controllers

import (
	"todo/src/db"
	"todo/src/models"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTodo(c fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	type body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status	  string `json:"status"`
	}

	var data body
if err := c.Bind().Body(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if data.Status != string(models.StatusCompleted) &&
		data.Status != string(models.StatusIncomplete) {
			data.Status = string(models.StatusIncomplete)
		}

	todo := bson.M{
		"_id": primitive.NewObjectID(),
		"title": data.Title,
		"description": data.Description,
		"status": data.Status,
		"userId": userId,
	}

	_, err := db.DB.Collection("todos").InsertOne(c.Context(), todo)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create todo",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{
			"message": "Todo created successfully",
			"todo": todo,
		},
	)
}

func GetTodos(c fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	cursor, err := db.DB.Collection("todos").Find(c.Context(), bson.M{
		"userId": userId,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot fetch todos",
		})
	}

	var todos []bson.M
	if err := cursor.All(c.Context(), &todos); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot parse todos",
		})
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"todos": todos,
		},
	)
}  // /:id

func DeleteTodo(c fiber.Ctx) error {
	todoId := c.Params("id")
	userId := c.Locals("userId").(string)

	objId, err := primitive.ObjectIDFromHex(todoId)   //string -> objectID
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid todo ID",
		})
	}

	filter := bson.M{
		"_id":    objId,
		"userId": userId,
	}

	result, err := db.DB.Collection("todos").DeleteOne(c.Context(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot delete todo",
		})
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo deleted successfully",	
	})
}

func UpdateTodo(c fiber.Ctx) error {
	todoId := c.Params("id")
	userId := c.Locals("userId").(string)

	objId, err := primitive.ObjectIDFromHex(todoId)   //string -> objectID
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid todo ID",
		})
	}

	type body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status	  string `json:"status"`
	}

	var data body
if err := c.Bind().Body(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	update := bson.M{}
	if data.Title != "" {
		update["title"] = data.Title
	}
	if data.Description != "" {
		update["description"] = data.Description
	}
	if data.Status != "" {
		update["status"] = data.Status
	}

	if len(update) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No fields to update",
		})
	}

	filter := bson.M{
		"_id":    objId,
		"userId": userId,
	}

	result, err := db.DB.Collection("todos").UpdateOne(c.Context(), filter, bson.M{
		"$set": update,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot update todo",
		})
	}

	if result.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo updated successfully",
	})
}