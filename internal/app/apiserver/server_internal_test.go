package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GuseynovAnar/rest_api.git/internal/app/store/sqlstore/storetestings/mocks"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	server := newServer(mocks.New())

	server.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}
