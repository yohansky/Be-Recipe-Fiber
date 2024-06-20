package models

type Liker struct {
	Id       uint `json:"id"`
	UserId   uint
	RecipeId uint
}
