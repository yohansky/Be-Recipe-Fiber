package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mod/src/config"
	"go.mod/src/models"
)

func AllComments(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(config.DB, &models.Comment{}, page))
}

func CreateComment(c *fiber.Ctx) error {
	var comment models.Comment

	if err := c.BodyParser(&comment); err != nil {
		return err
	}

	config.DB.Create(&comment)

	return c.JSON(comment)
}

func GetComment(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var comment models.Comment

	comment.Id = uint(id)

	config.DB.Preload("User").Find(&comment)

	return c.JSON(comment)
}

func GetCommentbyUserId(c *fiber.Ctx) error {
	id := c.Params("id")

	var saves []models.Comment
	if err := config.DB.Where("user_id = ?", id).Find(&saves).Error; err != nil {
		return c.JSON(fiber.Map{"error": "Saves not found"})
	}

	return c.JSON(saves)

}

func UpdateComment(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var saved models.Comment

	saved.Id = uint(id)

	if err := c.BodyParser(&saved); err != nil {
		return err
	}

	config.DB.Model(&saved).Updates(saved)

	return c.JSON(saved)
}

func DeleteComment(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var saved models.Comment

	saved.Id = uint(id)

	config.DB.Delete(&saved)

	return c.JSON(fiber.Map{
		"Message": "Delete Complete",
	})
}
