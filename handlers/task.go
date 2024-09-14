package handlers

import (
	"time"

	"github.com/Shaurya-Jha/real-time-task-manager/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// in memory data store (will replace this with a database later)
var storedTasks = make(map[string]models.Task)

// GetTasks return all tasks
func GetTasks(c *fiber.Ctx) error {
	// created tasksList array as per the Task model struct that stores all tasks
	tasksList := []models.Task{}

	// loop through storedTasks map / dictionary / json
	for _, task := range storedTasks {
		tasksList = append(tasksList, task)
	}

	// return the reponse as JSON
	return c.Status(fiber.StatusOK).JSON(tasksList)
}

// GetTask returns a task by ID
func GetTask(c *fiber.Ctx) error {
	// get id of task from parameter
	id := c.Params("id")

	task, exists := storedTasks[id]

	// check if the task exists or not as per the id from the storedTasks
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}

	// if task exists return in JSON format
	return c.Status(fiber.StatusOK).JSON(task)
}

// CreateTask creates a new task
func CreateTask(c *fiber.Ctx) error {
	// empty variable task as per the Task model struct
	var task models.Task

	// &task accesses the memory address of the task variable
	// when we create a variable, a memory address is allocated for the variable to store the value
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// add extra data to the fields for the task variable using 3rd party packages
	// create a new uuid for the id in the task
	task.ID = uuid.New().String()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	// store the created task in the storesTask variable as per the task id
	storedTasks[task.ID] = task

	return c.Status(fiber.StatusCreated).JSON(task)
}

// UpdateTask updates and existing task
func UpdateTask(c *fiber.Ctx) error {
	// get the for which we want to update the task from the params
	id := c.Params("id")
	// variable to temporarily store updated task and then send that as JSON
	var updatedTask models.Task

	if err := c.BodyParser(&updatedTask); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// store the tasks from the storedTasks as per the
	task, exists := storedTasks[id]

	// if tasks not exists that we want to update return Json response
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task does not exists",
		})
	}

	if updatedTask.Title != "" {
		task.Title = updatedTask.Title
	}

	if updatedTask.Description != "" {
		task.Description = updatedTask.Description
	}

	if !updatedTask.DueDate.IsZero() {
		task.DueDate = updatedTask.DueDate
	}

	if updatedTask.Status != "" {
		task.Status = updatedTask.Status
	}

	task.UpdatedAt = time.Now()
	// store the updated task in the storedTasks var
	storedTasks[id] = task

	return c.Status(fiber.StatusOK).JSON(task)
}

// DeleteTask deletes a task by ID
func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	_, exists := storedTasks[id]

	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}

	delete(storedTasks, id)

	return c.SendStatus(fiber.StatusNoContent)
}
