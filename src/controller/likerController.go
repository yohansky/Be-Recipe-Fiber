package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mod/src/config"
	"go.mod/src/models"
)

func Alllikes(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(config.DB, &models.Liker{}, page))
}

func CreateLikes(c *fiber.Ctx) error {
	var like models.Liker

	if err := c.BodyParser(&like); err != nil {
		return err
	}

	config.DB.Create(&like)

	return c.JSON(like)
}

func GetLiked(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var like models.Liker

	like.Id = uint(id)

	config.DB.Preload("Recipe").Find(&like)

	return c.JSON(like)
}

func GetLikebyUserId(c *fiber.Ctx) error {
	id := c.Params("id")

	var likes []models.Liker
	if err := config.DB.Where("user_id = ?", id).Find(&likes).Error; err != nil {
		return c.JSON(fiber.Map{"error": "Likes not found"})
	}

	return c.JSON(likes)

}

func UpdateLike(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var like models.Liker

	like.Id = uint(id)

	if err := c.BodyParser(&like); err != nil {
		return err
	}

	config.DB.Model(&like).Updates(like)

	return c.JSON(like)
}

func DeleteLike(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var like models.Liker

	like.Id = uint(id)

	config.DB.Delete(&like)

	return c.JSON(fiber.Map{
		"Message": "Delete Complete",
	})
}
