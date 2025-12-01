package models

type UserModel struct {
	Name string `gorm:"not null" json:"name"`
	Email    string `gorm:"unique;not null" json:"email" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
}