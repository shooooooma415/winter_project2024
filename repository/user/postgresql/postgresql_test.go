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
		{name: "userNameでユーザーを追加する", userName: "hoge"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			db, err := setupDB.ConnectDB()
			if err != nil {
				t.Fatalf("Failed to connect to the database: %v", err)
			}
			defer db.Close()

			repository := postgresql.NewUserRepository(db)
			got, err := repository.CreateUserQuery(tc.userName)
			want := tc.userName

			if err != nil {
				t.Fatalf("CreateUserQuery() error = %v", err)
			}
			if got.UserName != tc.userName {
				t.Errorf("postgresql.CreateUserQuery() = %v, want %v", got.UserName, want)
			}
		})
	}
}

func TestUpdateUserQuery(t *testing.T) {
	testCases := []struct {
		name     string
		userName model.UserName
		wantName model.UserName
	}{
		{name: "userNameをwantNameに変更する", userName: "hoge", wantName: "hogehoge"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			db, err := setupDB.ConnectDB()
			if err != nil {
				t.Fatalf("Failed to connect to the database: %v", err)
			}

			defer db.Close()
			repository := postgresql.NewUserRepository(db)

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
		})
	}
}
