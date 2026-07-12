package main

import (
	"log"

	"mesenshuttle-backend/internal/config"
	"mesenshuttle-backend/internal/controllers"
	"mesenshuttle-backend/internal/models"
	"mesenshuttle-backend/internal/repositories"
	"mesenshuttle-backend/internal/routes"
	"mesenshuttle-backend/internal/services"
	"mesenshuttle-backend/pkg/database"
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

	// Dependency Injection
	adminRepo := repositories.NewAdminRepository(db)
	authService := services.NewAuthService(adminRepo, cfg)
	authController := controllers.NewAuthController(authService)

	// Setup Router
	r := routes.SetupRouter(authController)

	log.Printf("Starting server on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
