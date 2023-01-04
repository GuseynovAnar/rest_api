package mocks

import (
	"github.com/GuseynovAnar/rest_api.git/internal/app/models"
	"github.com/GuseynovAnar/rest_api.git/internal/app/store"
)

// Store ...
type Store struct {
	repository *UserRepository
}

// New instance ...
func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.repository != nil {
		return s.repository
	}

	s.repository = &UserRepository{
		store: s,
		users: make(map[string]*models.User),
	}

	return s.repository
}
