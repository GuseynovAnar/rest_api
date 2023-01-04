package storetestings

import (
	"os"
	"testing"
)

var DatabaseURL string

// TestMain ...
func TestMain(m *testing.M) {
	DatabaseURL = os.Getenv("DATABASE_URL")

	if DatabaseURL == "" {
		DatabaseURL = "user=aguseynov password=23045 host=localhost dbname=restapi_test sslmode=disable"
	}

	os.Exit(m.Run())
}
