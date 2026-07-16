package main

import (
	"flag"
	"log"

	"mesenshuttle-backend/internal/config"
	"mesenshuttle-backend/internal/seeders"
	"mesenshuttle-backend/pkg/database"

	"github.com/pressly/goose/v3"
)

func main() {
	migrateFlag := flag.Bool("migrate", false, "Run database migrations (Up)")
	rollbackFlag := flag.Bool("rollback", false, "Rollback the last migration (Down)")
	statusFlag := flag.Bool("migrate-status", false, "Print migration status")
	seedFlag := flag.Bool("seed", false, "Run database seeders")
	flag.Parse()

	cfg := config.LoadConfig()

	db, err := database.InitMySQL(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get generic database object: %v", err)
	}

	if err := goose.SetDialect("mysql"); err != nil {
		log.Fatalf("Failed to set goose dialect: %v", err)
	}

	if *migrateFlag {
		log.Println("Running Migrations (Up)...")
		if err := goose.Up(sqlDB, "db/migrations"); err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}
		log.Println("Migrations completed successfully!")
	}

	if *rollbackFlag {
		log.Println("Rolling back last migration (Down)...")
		if err := goose.Down(sqlDB, "db/migrations"); err != nil {
			log.Fatalf("Failed to rollback migration: %v", err)
		}
		log.Println("Rollback completed successfully!")
	}

	if *statusFlag {
		log.Println("Migration Status:")
		if err := goose.Status(sqlDB, "db/migrations"); err != nil {
			log.Fatalf("Failed to get migration status: %v", err)
		}
	}

	if *seedFlag {
		log.Println("Running Seeders...")
		seeders.SeedAdmin(db, cfg)
		log.Println("Seeders completed successfully!")
	}

	if !*migrateFlag && !*rollbackFlag && !*statusFlag && !*seedFlag {
		log.Println("No flags provided. Use -migrate, -rollback, -migrate-status, or -seed.")
	}
}
