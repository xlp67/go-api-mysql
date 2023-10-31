package models

import (
	"api/database"
)

func Delete(user *database.User) (*database.User, error) {
	db, err := database.ConnectionDBUser()
	if err != nil {panic(err)}
	err = db.Where("User = ?", user.User).First(&user).Error 
	if err != nil {panic(err)}
		db.Delete(&user)

	return user, nil
}