package storetestings

import (
	"fmt"
	"strings"
	"testing"

	"github.com/GuseynovAnar/rest_api.git/internal/app/store"
)

// TestStore ...

func TestStore(t *testing.T, databaseURL string) (*store.Store, func(...string)) {
	t.Helper()

	config := store.NewConfig()
	config.DatabaseURL = databaseURL
	s := store.New(config)

	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.DB().Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}
