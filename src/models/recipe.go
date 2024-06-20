package models

import "gorm.io/gorm"

type Recipe struct {
	Id         uint `json:"id"`
	Name       string
	Ingredient string
	Photo      string
	VideoUrl   string
	UserId     uint
}

func (recipe *Recipe) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Recipe{}).Count(&total)

	return total
}

func (recipe *Recipe) Take(db *gorm.DB, limit int, offset int) interface{} {
	var recipes []Recipe

	db.Offset(offset).Limit(limit).Find(&recipes)

	return recipes
}
