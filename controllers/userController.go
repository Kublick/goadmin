package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kublick/goadmin/database"
	"github.com/kublick/goadmin/models"
)

func AllUsers(c *fiber.Ctx) error {

	var users []models.User

	database.DB.Find(&users)

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Create(&user)

	user.SetPassword("1234")

	database.DB.Save(&user)

	return c.JSON(user)

}
