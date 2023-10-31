package database

import (
	"api/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type User struct {
	ID int
	User string
	Password string
}

func ConnectionDBUser() (*gorm.DB, error) {
	var user *User
	conf, err := config.LoadConfig(".")
	if err != nil {panic(err)}
	credential := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName)
		db, err := gorm.Open(mysql.Open(credential), &gorm.Config{})
		if err != nil {panic(err)}
		db.AutoMigrate(&user)

	return db, nil
}