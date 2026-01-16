package db

import (
	"github.com/todo-app/internal/config"
	"github.com/todo-app/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() {
	var err error
	switch config.AppCfg.DbConfig.Driver {
	case "sqlite":
		DB, err = gorm.Open(sqlite.Open(config.AppCfg.DbConfig.Dsn), &gorm.Config{})
	default:
		log.Fatalf("unsupported db driver: %s", config.AppCfg.DbConfig.Driver)
	}

	if err != nil {
		log.Fatal("failed to connect database", err)
	}
}

func MigrateAll() {
	err := DB.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		log.Fatal("failed to AutoMigrate ")
	}
}
