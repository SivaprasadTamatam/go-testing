package dbops

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	SetupDB()
	defer TeardownDB()

	code := m.Run()

	os.Exit(code)
}

func TestQueryDatabase(t *testing.T) {
	db.Query("Select * from table Customer")
}
