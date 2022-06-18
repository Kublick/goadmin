package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kublick/goadmin/database"
	"github.com/kublick/goadmin/models"
)

func AllProducts(c *fiber.Ctx) error {

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 2
	offset := (page - 1) * limit

	var total int64

	var products []models.Product

	database.DB.Offset(offset).Limit(limit).Find(&products)

	database.DB.Model(&models.Product{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": products,
		"meta": fiber.Map{
			"page":      page,
			"limit":     limit,
			"total":     total,
			"last_page": float64(int(total) / limit),
		},
	})
}

func CreateProduct(c *fiber.Ctx) error {

	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		fmt.Println("hay un error", err)
		return err
	}
	database.DB.Create(&product)

	return c.JSON(product)

}

func GetProduct(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{Id: uint(id)}

	database.DB.Find(&product)

	return c.JSON(product)

}

func UpdateProduct(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{Id: uint(id)}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(product)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{Id: uint(id)}

	database.DB.Delete(&product)

	return c.JSON(fiber.Map{
		"message": "Product deleted",
	})
}
