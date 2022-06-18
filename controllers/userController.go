package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kublick/goadmin/database"
	"github.com/kublick/goadmin/models"
)

func AllUsers(c *fiber.Ctx) error {

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 2
	offset := (page - 1) * limit

	var total int64

	var users []models.User

	database.DB.Preload("Role").Offset(offset).Limit(limit).Find(&users)

	database.DB.Model(&models.User{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": users,
		"meta": fiber.Map{
			"page":      page,
			"limit":     limit,
			"total":     total,
			"last_page": float64(int(total) / limit),
		},
	})
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

func GetUser(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{Id: uint(id)}

	database.DB.Preload("Role").Find(&user)

	return c.JSON(user)

}

func UpdateUser(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{Id: uint(id)}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{Id: uint(id)}

	database.DB.Delete(&user)

	return c.JSON(fiber.Map{
		"message": "User deleted",
	})
}
