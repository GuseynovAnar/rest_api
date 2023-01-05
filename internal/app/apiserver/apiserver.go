package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/GuseynovAnar/rest_api.git/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
)

// Start ...
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)

	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	server := newServer(store, sessionStore)

	return http.ListenAndServe(config.BindAdd, server)
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
