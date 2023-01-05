package mocks

import (
	"github.com/GuseynovAnar/rest_api.git/internal/app/models"
	"github.com/GuseynovAnar/rest_api.git/internal/app/store"
)

// Mock UserRepository ...
type UserRepository struct {
	store *Store
	users map[int]*models.User
}

// Mock Create ...
func (r *UserRepository) Create(model *models.User) error {

	if err := model.Validate(); err != nil {
		return err
	}

	if err := model.PreCreatePhase(); err != nil {
		return err
	}

	model.ID = len(r.users) + 1
	r.users[model.ID] = model

	return nil
}

// Mock FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFoud
}

func (r *UserRepository) Find(id int) (*models.User, error) {

	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFoud
	}

	return u, nil
}
