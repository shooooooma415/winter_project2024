package postgresql_test

import (
	"testing"
	"winter_pj/model"
	postgresql "winter_pj/repository/user/postgresql"
	setupDB "winter_pj/repository/utils"

	_ "github.com/lib/pq"
)

func TestCreateUserQuery(t *testing.T) {
	testCases := []struct {
		name     string
		userName model.UserName
	}{
		{name: "Valid", userName: "hoge"},
		{name: "Empty", userName: ""},
	}

	db, err := setupDB.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
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
		name     string
		userName model.UserName
		wantName model.UserName
	}{
		{name: "Valid", userName: "hoge", wantName: "hogehoge"},
	}

	db, err := setupDB.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	defer db.Close()
	repository := postgresql.NewUserRepository(db)

	for _, tc := range testCases {
		user, err := repository.CreateUserQuery(tc.userName)
		if err != nil {
			t.Fatalf("CreateUserQuery() error = %v", err)
		}

		updateUser := model.User{
			UserId:   user.UserId,
			UserName: tc.wantName,
		}
		got, err := repository.UpdateUserQuery(updateUser)
		want := updateUser

		if err != nil {
			t.Fatalf("CreateUserQuery() error = %v", err)
		}
		if got.UserName != want.UserName {
			t.Errorf("postgresql.UpdateUserQuery() = %v, want %v", got, want)
		}
	}
}
