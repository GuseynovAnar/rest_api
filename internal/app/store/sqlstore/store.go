package sqlstore

import (
	"database/sql"

	"github.com/GuseynovAnar/rest_api.git/internal/app/store"
	_ "github.com/lib/pq" // ...
)

// Store ...
type Store struct {
	db         *sql.DB
	repository *UserRepository
}

// New instance ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.repository != nil {
		return s.repository
	}

	s.repository = &UserRepository{
		store: s,
	}

	return s.repository
}
