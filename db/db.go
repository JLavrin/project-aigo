package db

import (
	"fmt"
	"github.com/JLavrin/project-aigo/config"
	"github.com/JLavrin/project-aigo/src/domain/article/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

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

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	migrationErr := DB.AutoMigrate(&model.Article{})

	if migrationErr != nil {
		log.Fatalf("models.Setup err: %v", migrationErr)
	}

	log.Println("Database connection established.")
	DB.Config.Logger.LogMode(3)
}
