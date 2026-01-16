package repository

import (
	"github.com/todo-app/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByProviderId(provider, providerId string) (*models.User, error) {
	var user models.User
	err := r.db.Where("provider = ? AND provider_id = ?", provider, providerId).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(entity *models.User) error {
	return r.db.Create(entity).Error
}
