package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 100)),
	)
}

func (user *User) PreCreatePhase() error {

	if len(user.Password) > 0 {
		encrypted, err := encryptString(user.Password)

		if err != nil {
			return err
		}

		user.EncryptedPassword = encrypted
	}

	return nil
}

func (user *User) Sanitize() {
	user.Password = ""
}

func (user *User) ComparePassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(pass)) == nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(b), nil
}
