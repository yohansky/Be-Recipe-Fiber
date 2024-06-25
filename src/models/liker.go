package models

import "gorm.io/gorm"

type Liker struct {
	Id       uint `json:"id"`
	UserId   uint
	RecipeId uint
	Recipe   Recipe `gorm:"foreignKey:RecipeId"`
}

func (liker *Liker) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Liker{}).Count(&total)

	return total
}

func (liker *Liker) Take(db *gorm.DB, limit int, offset int) interface{} {
	var likes []Liker

	db.Offset(offset).Limit(limit).Find(&likes)

	return likes
}
