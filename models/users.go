package models

import "time"

type User struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Email string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (User) GetTableName() string {
	return "users"
}