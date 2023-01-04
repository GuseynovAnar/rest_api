package store

import "github.com/GuseynovAnar/rest_api.git/internal/app/models"

// UserRepository ...
type UserRepository interface {
	Create(model *models.User) error
	FindByEmail(email string) (*models.User, error)
}
