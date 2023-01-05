package sqlstoretestings_test

import (
	"testing"

	"github.com/GuseynovAnar/rest_api.git/internal/app/models"
	"github.com/GuseynovAnar/rest_api.git/internal/app/store"
	"github.com/GuseynovAnar/rest_api.git/internal/app/store/sqlstore"
	sqlstoretestings "github.com/GuseynovAnar/rest_api.git/internal/app/store/sqlstore/storetestings"
	"github.com/GuseynovAnar/rest_api.git/internal/app/store/sqlstore/storetestings/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateWithDB(t *testing.T) {
	db, teardown := sqlstoretestings.TestDB(t, sqlstoretestings.DatabaseURL)

	defer teardown("users")

	s := sqlstore.New(db)
	u := models.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindeByEmailWithDB(t *testing.T) {
	db, teardown := sqlstoretestings.TestDB(t, sqlstoretestings.DatabaseURL)

	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@example.com"

	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFoud.Error())

	u := models.TestUser(t)
	u.Email = email

	s.User().Create(u)

	user, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_FindeByIdWithDB(t *testing.T) {
	db, teardown := sqlstoretestings.TestDB(t, sqlstoretestings.DatabaseURL)

	defer teardown("users")

	s := sqlstore.New(db)
	id := 1

	_, err := s.User().Find(id)
	assert.EqualError(t, err, store.ErrRecordNotFoud.Error())

	u := models.TestUser(t)
	u.ID = id

	s.User().Create(u)

	user, err := s.User().Find(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_CreateMock(t *testing.T) {

	s := mocks.New()
	u := models.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindeByEmailMock(t *testing.T) {
	s := mocks.New()
	email := "user@example.com"

	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFoud.Error())

	u := models.TestUser(t)
	u.Email = email

	s.User().Create(u)

	user, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_FindeByIDMock(t *testing.T) {
	s := mocks.New()
	id := 1

	_, err := s.User().Find(id)
	assert.EqualError(t, err, store.ErrRecordNotFoud.Error())

	u := models.TestUser(t)
	u.ID = id

	s.User().Create(u)

	user, err := s.User().Find(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
