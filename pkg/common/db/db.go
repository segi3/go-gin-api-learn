package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"segi3.com/api/pkg/common/models"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
