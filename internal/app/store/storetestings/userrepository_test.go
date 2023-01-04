package storetestings_test

import (
	"testing"

	"github.com/GuseynovAnar/rest_api.git/internal/app/models"
	_ "github.com/GuseynovAnar/rest_api.git/internal/app/store"
	"github.com/GuseynovAnar/rest_api.git/internal/app/store/storetestings"

	"github.com/stretchr/testify/assert"
)

func TestUserReposit_Create(t *testing.T) {
	store, teardown := storetestings.TestStore(t, storetestings.DatabaseURL)

	defer teardown("users")

	user, err := store.GetUser().Create(models.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserReposit_FindeByEmail(t *testing.T) {
	store, teardown := storetestings.TestStore(t, storetestings.DatabaseURL)

	defer teardown("users")

	email := "user@example.com"

	_, err := store.GetUser().FindByEmail(email)
	assert.Error(t, err)

	u := models.TestUser(t)
	u.Email = email

	store.GetUser().Create(u)

	user, err := store.GetUser().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
