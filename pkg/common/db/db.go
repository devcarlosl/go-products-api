package db

import (
	"fmt"
	"github.com/devcarlosl/go-products-api/pkg/common/config"
	"github.com/devcarlosl/go-products-api/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init(c *config.Config) *gorm.DB {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Product{})

	return db
}
