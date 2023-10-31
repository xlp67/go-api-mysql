package models

import (
	"api/database"
	"log"
)

func UpdateUser(data *database.User) error {
	var validate *database.User
	db, err := database.ConnectionDBUser()
	if err != nil {panic(err)}
	err = db.Where("user = ?", data.User).Find(&validate).Error
	if err != nil {
		log.Println("User not found") 
	return err
}
	if data.User == validate.User {
		if data.Password != validate.Password {
			db.Delete(&validate)
			validate = &database.User{
				ID: validate.ID,
				User: data.User,
				Password: data.Password,
			}
			db.Create(&validate)
			} else {log.Println("Equals passwords")}
	} else {log.Println("User not found")}
	return nil
}