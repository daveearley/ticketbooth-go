package database

import (
	"testing"
)

func TestInitDb(t *testing.T) {
	db := InitDb()

	if db.Close() != nil {
		t.Error("Failed to load & close db connection")
	}
}
