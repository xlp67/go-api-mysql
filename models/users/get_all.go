package models

import (
	"api/database"
)

func GetAll()  ([]database.User,error) {
	var datas []database.User
	db, err := database.ConnectionDBUser()
	if err != nil {panic(err)}
	db.Find(&datas)
return  datas, err
}