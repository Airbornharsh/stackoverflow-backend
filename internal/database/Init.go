package database

import "github.com/joho/godotenv"

func INIT() {
	dotenv := godotenv.Load(".env")

	if dotenv != nil {
		panic("Error loading .env file")
	}

	// Connect to database
	// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// defer db.Close()
}
