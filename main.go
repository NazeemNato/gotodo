package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nazeemnato/gotodo/database"
	"github.com/nazeemnato/gotodo/handler"
	"github.com/nazeemnato/gotodo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() {
	var err error
	dsn := "host=localhost user=postgres password=1234567890 dbname=gotodo port=5432"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected")
	database.DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("Database migrated")
}

func setupRoutes(app *fiber.App) {
	app.Get("/todo", handler.GetTodos)
	app.Post("/todo", handler.CreateTodo)
	app.Get("/todo/:id", handler.GetTodoById)

}

func main() {
	app := fiber.New()
	initDb()
	setupRoutes(app)
	app.Listen(":8000")
}
