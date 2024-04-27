package models

import (
	"fmt"
	"github.com/JLavrin/project-aigo/config"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

var db *gorm.DB

func Setup() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Database.Host,
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
		config.Database.Port,
		config.Database.SSLMode,
		config.Database.TimeZone,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	migrationErr := db.AutoMigrate(&Article{})

	if migrationErr != nil {
		log.Fatalf("models.Setup err: %v", migrationErr)
	}
	log.Default().Println("Database connection established")
	db.Config.Logger.LogMode(3)
}
