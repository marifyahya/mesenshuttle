package main

import (
	"log"

	"mesenshuttle-backend/internal/config"
	"mesenshuttle-backend/internal/models"
	"mesenshuttle-backend/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.InitMySQL(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize MySQL: %v", err)
	}

	log.Println("Running Auto Migration...")
	err = db.AutoMigrate(
		&models.Admin{},
		&models.Route{},
		&models.Fleet{},
		&models.Schedule{},
		&models.Booking{},
		&models.BookingSeat{},
	)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Auto Migration completed!")

	_, err = database.InitRedis(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Mesenshuttle Backend is running!",
		})
	})

	log.Printf("Starting server on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
