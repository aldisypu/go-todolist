package controller

import (
	"strconv"

	"github.com/aldisypu/go-todolist/config"
	"github.com/aldisypu/go-todolist/models"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	config.DB.Create(&todo)

	return c.Status(200).JSON(todo)
}

func FindAll(c *fiber.Ctx) error {
	todos := []models.Todo{}

	config.DB.Find(&todos)

	return c.Status(200).JSON(todos)
}

func FindById(c *fiber.Ctx) error {
	todos := []models.Todo{}

	config.DB.First(&todos, c.Params("todoId"))

	return c.Status(200).JSON(todos)
}

func Update(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	todoId, _ := strconv.Atoi(c.Params("todoId"))

	config.DB.Model(&models.Todo{}).Where("id = ?", todoId).Updates(&todo)

	return c.Status(200).JSON("Success update todo")
}

func Delete(c *fiber.Ctx) error {
	todo := new(models.Todo)

	todoId, _ := strconv.Atoi(c.Params("todoId"))

	config.DB.Where("id = ?", todoId).Delete(&todo)

	return c.Status(200).JSON("Success delete Todo")
}
