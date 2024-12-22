package setupDB_test

import (
	"testing"
	setupDB "winter_pj/repository/utils"

	_ "github.com/lib/pq"
)

func TestConnectDB(t *testing.T) {
	db, err := setupDB.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	if db == nil {
		t.Fatal("Expected a valid database connection, but got nil")
	}
	defer db.Close()
}
