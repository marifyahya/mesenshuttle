package database

import (
	"fmt"
	"log"

	"mesenshuttle-backend/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to MySQL: %v", err)
		return nil, err
	}

	log.Println("Connected to MySQL successfully!")
	return db, nil
}
