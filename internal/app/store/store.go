package store

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
)

// Store ...
type Store struct {
	config *Config
	db     *sql.DB
}

// New instance ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open store ...

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close store ...
func (s *Store) Close() {
	s.db.Close()
}
