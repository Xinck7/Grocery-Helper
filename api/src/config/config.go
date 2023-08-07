package config

import (
	"fmt"
	"os"
	"xinck/api/src/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error loading .env file")
	}

	var dbhost, dbuser, dbpass, dbname, dbport = os.Getenv("host"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("port")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbhost, dbuser, dbpass, dbname, dbport)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// log.Fatal("unable to connect to database")
		panic("failed to connect to database")
	}
	db.Debug().AutoMigrate(&models.User{}, &models.Ingredient{}, &models.Item{}, &models.List{}, &models.Recipe{})
	return db
}

func DisconnectDB(db *gorm.DB) {
	dbPostgres, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	dbPostgres.Close()
}
