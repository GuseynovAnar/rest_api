package mocks

import (
	"github.com/GuseynovAnar/rest_api.git/internal/app/models"
	"github.com/GuseynovAnar/rest_api.git/internal/app/store"
)

// Mock UserRepository ...
type UserRepository struct {
	store *Store
	users map[string]*models.User
}

// Mock Create ...
func (r *UserRepository) Create(model *models.User) error {

	if err := model.Validate(); err != nil {
		return err
	}

	if err := model.PreCreatePhase(); err != nil {
		return err
	}

	r.users[model.Email] = model

	model.ID = len(r.users)

	return nil
}

// Mock FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {

	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrRecordNotFoud
	}

	return u, nil
}
