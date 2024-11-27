package database

import (
	"fmt"
	"log"
	"strconv"

	"gofiber-boiler/config"
	"gofiber-boiler/internals/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Fail to connect to DB")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))

	fmt.Println("dsn")
	fmt.Println(dsn)
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		fmt.Println(dsn)
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// Migrate the database
	DB.AutoMigrate(&model.Post{})
	DB.AutoMigrate(&model.User{})
	fmt.Println("Database Migrated")
}
