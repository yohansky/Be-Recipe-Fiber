package models

import "gorm.io/gorm"

type Comment struct {
	Id       uint `json:"id"`
	Text     string
	UserId   uint
	User     User `gorm:"foreignKey:UserId"`
	RecipeId uint
}

func (comment *Comment) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Comment{}).Count(&total)

	return total
}

func (comment *Comment) Take(db *gorm.DB, limit int, offset int) interface{} {
	var comments []Comment

	db.Offset(offset).Limit(limit).Find(&comments)

	return comments
}
