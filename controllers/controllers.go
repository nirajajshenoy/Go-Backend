package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// "errors"

func Controller() error {
	return errors.New("an error occurred")
	// fmt.Println("WELCOME")make

}
func GetController(c *fiber.Ctx) error {

	return c.SendString("Get Data \n")
}
func PostController(c *fiber.Ctx) error {

	return c.SendString("Post Data \n")
}
func DeleteController(c *fiber.Ctx) error {

	return c.SendString("Delete Data \n")
}