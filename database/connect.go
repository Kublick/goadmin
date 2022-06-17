package database

import (
	"fmt"

	"github.com/kublick/goadmin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	database, err := gorm.Open(mysql.Open("root:0QRIyeyOmwotb6aS7oUN@tcp(containers-us-west-61.railway.app:7959)/railway"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.User{}, &models.Role{})

	DB = database

	fmt.Println("Connected to database")

}
