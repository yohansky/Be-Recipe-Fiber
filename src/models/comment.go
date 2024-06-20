package models

type Comment struct {
	Id       uint `json:"id"`
	Text     string
	UserId   uint
	RecipeId uint
}
