package sqlstore

import (
	"database/sql"

	"github.com/GuseynovAnar/rest_api.git/internal/app/models"
	"github.com/GuseynovAnar/rest_api.git/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(model *models.User) error {

	if err := model.Validate(); err != nil {
		return err
	}

	if err := model.PreCreatePhase(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		model.Email,
		model.EncryptedPassword,
	).Scan(&model.ID)
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	u := &models.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFoud
		}
		return nil, err
	}

	return u, nil
}
