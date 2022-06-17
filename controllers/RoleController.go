package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kublick/goadmin/database"
	"github.com/kublick/goadmin/models"
)

func AllRoles(c *fiber.Ctx) error {

	var roles []models.Role

	database.DB.Preload("Permissions").Find(&roles)

	return c.JSON(roles)
}

func CreateRole(c *fiber.Ctx) error {
	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	database.DB.Save(&role)

	return c.JSON(role)

}

func GetRole(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{Id: uint(id)}

	database.DB.Preload("Permissions").Find(&role)

	return c.JSON(role)

}

func UpdateRole(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	roleDto := fiber.Map{}

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	database.DB.Exec("DELETE FROM role_permissions WHERE role_id = ?", id)

	role := models.Role{
		Id:          uint(id),
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	database.DB.Model(&role).Updates(role)

	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{Id: uint(id)}

	database.DB.Delete(&role)

	return c.JSON(fiber.Map{
		"message": "Role deleted",
	})
}
