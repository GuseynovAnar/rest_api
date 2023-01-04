package models_test

import (
	"testing"

	"github.com/GuseynovAnar/rest_api.git/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := models.TestUser(t)

	assert.NoError(t, u.PreCreatePhase())
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *models.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *models.User {
				return models.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Email = "invalidemail"
				return u
			},
			isValid: false,
		},
		{
			name: "invalid password lenght 1",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Password = "pass"
				return u
			},
			isValid: false,
		},
		{
			name: "valid password lenght",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Password = "123456"
				return u
			},
			isValid: true,
		},
		{
			name: "empty password",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "has encrypted_password and valid",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "EncryptedPassword"
				return u
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				if tc.isValid {
					assert.NoError(t, tc.u().Validate())
				} else {
					assert.Error(t, tc.u().Validate())
				}
			},
		)
	}
}
