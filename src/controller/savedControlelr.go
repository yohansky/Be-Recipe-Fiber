package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mod/src/config"
	"go.mod/src/models"
)

func AllSaves(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(config.DB, &models.Saved{}, page))
}

func CreateSave(c *fiber.Ctx) error {
	var saved models.Saved

	if err := c.BodyParser(&saved); err != nil {
		return err
	}

	config.DB.Create(&saved)

	return c.JSON(saved)
}

func GetSave(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var saved models.Saved

	saved.Id = uint(id)

	config.DB.Preload("Recipe").Find(&saved)

	return c.JSON(saved)
}

func GetSavesbyUserId(c *fiber.Ctx) error {
	id := c.Params("id")

	var saves []models.Saved
	if err := config.DB.Where("user_id = ?", id).Find(&saves).Error; err != nil {
		return c.JSON(fiber.Map{"error": "Saves not found"})
	}

	return c.JSON(saves)

}

func UpdateSave(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var saved models.Saved

	saved.Id = uint(id)

	if err := c.BodyParser(&saved); err != nil {
		return err
	}

	config.DB.Model(&saved).Updates(saved)

	return c.JSON(saved)
}

func DeleteSave(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var saved models.Saved

	saved.Id = uint(id)

	config.DB.Delete(&saved)

	return c.JSON(fiber.Map{
		"Message": "Delete Complete",
	})
}
