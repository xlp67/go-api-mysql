package models

import (
	"api/database"
)


func GetUser(user *database.User) (*database.User, error) {
	db, err := database.ConnectionDBUser()
	if err != nil {panic(err)}
	db.Where("user = ?", user.User).Find(&user)
	return user, nil
}