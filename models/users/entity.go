package models

type User struct {
	ID       string `gorm:"primaryKey"`
	User     string `json:"user"`
	Password  string `json:"-"`
}


