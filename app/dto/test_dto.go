package dto

type TestDto struct {
	Id int `form:"id" json:"id" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
}