package models

type User struct {
	UserID   uint   `gorm:"primaryKey" json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	UserType string `json:"user_type" gorm:"default:'user'"`
}
