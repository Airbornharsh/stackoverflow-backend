package database

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func INIT() {
	dotenv := godotenv.Load(".env")

	if dotenv != nil {
		panic("Error loading .env file")
	}

	DB_URI := os.Getenv("DB_URI")

	sqlDB, err := sql.Open("postgres", DB_URI)
	if err != nil {
		panic("failed to connect database")
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = gormDB

	// TableCreate()
}
