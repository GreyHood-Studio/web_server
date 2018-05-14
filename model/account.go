package model

type Account struct {
	ID     string `form:"id" json:"id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type User struct {
	ID     string `form:"id" json:"id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	NickName string `form:"nickname" json:"nickname" binding:"required"`
}