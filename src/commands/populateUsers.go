package main

import (
	"log"

	"github.com/liridonrama/microservices-with-go-users/src/database"
	"github.com/liridonrama/microservices-with-go-users/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	db, err := gorm.Open(mysql.Open("root:root@tcp(host.docker.internal:33066)/ambassador"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error while trying to run populateUsers.go:", err)
	}

	var users []models.User

	db.Find(&users)

	for _, user := range users {
		database.DB.Create(&user)
	}
}
