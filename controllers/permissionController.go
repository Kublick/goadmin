package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kublick/goadmin/database"
	"github.com/kublick/goadmin/models"
)

func AllPermissions(c *fiber.Ctx) error {

	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
