package helper

import (
	"go.mod/src/config"
	"go.mod/src/models"
)

func Migrate() {
	config.DB.AutoMigrate(
		&models.User{},
		&models.Comment{},
		&models.Liker{},
		&models.Recipe{},
		&models.Saved{},
	)
}
