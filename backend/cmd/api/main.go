package main

import (
	"log"

	"mesenshuttle-backend/internal/config"
	"mesenshuttle-backend/internal/container"
	"mesenshuttle-backend/internal/routes"
	"mesenshuttle-backend/pkg/database"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.InitMySQL(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize MySQL: %v", err)
	}

	redisClient, err := database.InitRedis(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}

	authModule := container.NewAuthModule(db, cfg, redisClient)
	routeModule := container.NewRouteModule(db)
	fleetModule := container.NewFleetModule(db)
	scheduleModule := container.NewScheduleModule(db)

	modules := []container.Module{
		authModule,
		routeModule,
		fleetModule,
		scheduleModule,
	}

	r := routes.SetupRouter(modules, cfg.JWTSecret, redisClient)

	log.Printf("Starting server on port %s...", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
