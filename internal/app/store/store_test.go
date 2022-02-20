package store

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgresql://root:123@localhost:5432/restapi_dev?sslmode=disable"
	}

	os.Exit(m.Run())
}
