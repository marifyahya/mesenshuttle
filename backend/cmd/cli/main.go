package main

import (
	"flag"
	"log"

	"mesenshuttle-backend/internal/config"
	"mesenshuttle-backend/internal/models"
	"mesenshuttle-backend/internal/seeders"
	"mesenshuttle-backend/pkg/database"
)

func main() {
	migrateFlag := flag.Bool("migrate", false, "Run database migrations")
	seedFlag := flag.Bool("seed", false, "Run database seeders")
	flag.Parse()

	cfg := config.LoadConfig()

	db, err := database.InitMySQL(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if *migrateFlag {
		log.Println("Running Auto Migration...")
		err = db.AutoMigrate(
			&models.Admin{},
			&models.Fleet{},
			&models.Route{},
			&models.Schedule{},
			&models.Booking{},
			&models.BookingSeat{},
		)
		if err != nil {
			log.Fatalf("Failed to auto migrate: %v", err)
		}
		log.Println("Auto Migration completed successfully!")
	}

	if *seedFlag {
		log.Println("Running Seeders...")
		seeders.SeedAdmin(db, cfg)
		log.Println("Seeders completed successfully!")
	}

	if !*migrateFlag && !*seedFlag {
		log.Println("No flags provided. Use -migrate or -seed.")
	}
}
