package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Create a Config struct with the desired options
	config := gorm.Config{
		Dialector: mysql.New(mysql.Config{
			// Set the username and password
			DSN: "customuser:custompassword@tcp(localhost:3308)/mydatabase?charset=utf8&parseTime=True&loc=Local"}),
	}

	// Open the database connection using the Config struct
	db, err := gorm.Open(config)
	if err != nil {
		panic("failed to connect database")
	}
	//db.AutoMigrate(User{})

	users := GenerateUsers(50)
	db.Create(&users)
	//db.Updates(&user)
	//db.Delete(&user)
}
