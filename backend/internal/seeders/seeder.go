package seeders

import (
	"log"

	"mesenshuttle-backend/internal/models"
	"mesenshuttle-backend/pkg/utils"

	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	var count int64
	db.Model(&models.Admin{}).Count(&count)
	if count == 0 {
		hashedPassword, err := utils.HashPassword("password123")
		if err != nil {
			log.Fatalf("Failed to hash password for admin seeder: %v", err)
		}

		admin := models.Admin{
			Name:         "Super Admin",
			Email:        "admin@example.com",
			PasswordHash: hashedPassword,
		}

		if err := db.Create(&admin).Error; err != nil {
			log.Fatalf("Failed to seed admin: %v", err)
		}
		log.Println("Admin seeded successfully!")
	} else {
		log.Println("Admin already exists, skipping seed.")
	}
}
