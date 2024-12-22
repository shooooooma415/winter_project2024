package postgresql_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"winter_pj/model"
	"winter_pj/repository/user/postgresql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func setupDB(t *testing.T) *sql.DB {
	if err := godotenv.Load("../../../.env"); err != nil {
		t.Fatalf("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		t.Fatalf("Failed to ping the database: %v", err)
	}
	return db
}

func TestCreateUserQuery(t *testing.T) {
	testCases := []struct {
		name     string
		userName model.UserName
	}{
		{name: "Valid", userName: "hoge"},
		{name: "Empty", userName: ""},
	}

	db := setupDB(t)
	defer db.Close()
	repository := postgresql.NewUserRepository(db)

	for _, tc := range testCases {
		got, err := repository.CreateUserQuery(tc.userName)
		want := tc.userName

		if err != nil {
			t.Fatalf("CreateUserQuery() error = %v", err)
		}
		if got.UserName != tc.userName {
			t.Errorf("postgresql.CreateUserQuery() = %v, want %v", got.UserName, want)
		}
	}
}

func TestUpdateUserQuery(t *testing.T) {
	testCases := []struct {
		name        string
		initialName model.UserName
		updateName  model.UserName
	}{
		{name: "Valid", initialName: "hoge", updateName: "hogehoge"},
		{name: "Empty", initialName: "hogehoge", updateName: ""},
	}

	db := setupDB(t)
	defer db.Close()
	repository := postgresql.NewUserRepository(db)

	for _, tc := range testCases {
		user, err := repository.CreateUserQuery(tc.initialName)
		if err != nil {
			t.Fatalf("CreateUserQuery() error = %v", err)
		}

		updateUser := model.User{
			UserId:   user.UserId,
			UserName: tc.updateName,
		}
		got, err := repository.UpdateUserQuery(updateUser)
		want := updateUser

		if err != nil {
			t.Fatalf("CreateUserQuery() error = %v", err)
		}
		if got.UserName != tc.updateName {
			t.Errorf("postgresql.UpdateUserQuery() = %v, want %v", got, want)
		}
	}
}
