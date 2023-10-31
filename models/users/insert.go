package models

import (
	"golang.org/x/crypto/bcrypt"
	"api/database"
)

func InsertUser(data *database.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {panic(err)}
	db, err := database.ConnectionDBUser()
	if err != nil {panic(err)}
	err = db.Where("user = ?", data.User).First(&data).Error 
	if err != nil {
		data = &database.User{
			User: data.User,
			Password: string(hash),
		}
		db.Create(&data)
	}
	return nil
}


