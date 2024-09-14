package main

import (
	"fmt"

	"github.com/Shaurya-Jha/real-time-task-manager/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// routes for task actions
	app.Get("/api/tasks", handlers.GetTasks)
	app.Get("/api/tasks/:id", handlers.GetTask)
	app.Post("/api/tasks", handlers.CreateTask)
	app.Put("/api/tasks/:id", handlers.UpdateTask)
	app.Delete("/api/tasks/:id", handlers.DeleteTask)

	// start the server
	app.Listen(":5000")
	fmt.Println("Server started on port: 5000")
}
