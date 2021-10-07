package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nazeemnato/gotodo/database"
	"github.com/nazeemnato/gotodo/models"
)

func GetTodos(c *fiber.Ctx) error {
	db := database.DBConn
	var todos []models.Todo

	db.Find(&todos)
	return c.JSON(&todos)
}

func CreateTodo(c *fiber.Ctx) error {
	db := database.DBConn
	todo := new(models.Todo)
	err := c.BodyParser(todo)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Invalid body", "error": err})
	}

	err = db.Create(&todo).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "something went wrong", "error": err})
	}

	return c.JSON(fiber.Map{"message": "Todo added"})
}

func GetTodoById(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var todo models.Todo
	err := db.Find(&todo, id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Unable to find todo", "error": err})
	}

	return c.JSON(&todo)
}