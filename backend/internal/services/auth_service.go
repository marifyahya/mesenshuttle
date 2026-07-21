package services

import (
	"errors"

	"gorm.io/gorm"
	"mesenshuttle-backend/internal/config"
	"mesenshuttle-backend/internal/models"
	"mesenshuttle-backend/internal/repositories"
	"mesenshuttle-backend/pkg/apperrors"
	"mesenshuttle-backend/pkg/utils"
)

type AuthService interface {
	Login(email, password string) (string, *models.Admin, error)
}

type authService struct {
	adminRepo repositories.AdminRepository
	cfg       *config.Config
}

func NewAuthService(adminRepo repositories.AdminRepository, cfg *config.Config) AuthService {
	return &authService{
		adminRepo: adminRepo,
		cfg:       cfg,
	}
}

func (s *authService) Login(email, password string) (string, *models.Admin, error) {
	admin, err := s.adminRepo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, apperrors.NewUnauthorized("invalid email or password")
		}
		return "", nil, err
	}

	if !utils.CheckPasswordHash(password, admin.PasswordHash) {
		return "", nil, apperrors.NewUnauthorized("invalid email or password")
	}

	secret := s.cfg.JWTSecret
	if secret == "" {
		return "", nil, apperrors.New(500, "JWT_SECRET is not configured")
	}

	token, err := utils.GenerateToken(admin.ID, admin.Email, secret)
	if err != nil {
		return "", nil, apperrors.New(500, "failed to generate token")
	}

	return token, admin, nil
}
