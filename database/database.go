package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-blognews/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	database := os.Getenv("PGDATABASE")
	sslmode := os.Getenv("PGSSLMODE")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta", host, user, password, database, port, sslmode)
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database", err)
	}

	db.Debug().AutoMigrate(
		&models.Users{},
		&models.Roles{},
		&models.Permissions{},
		&models.RolePermissions{},
		&models.Categories{},
		&models.Tags{},
		&models.Articles{},
		&models.ArticleTags{},
		&models.Likes{},
		&models.Comments{},
	)
}

func GetDB() *gorm.DB {
	return db
}
