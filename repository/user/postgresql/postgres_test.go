package postgresql

import (
	"testing"
	"winter_pj/model"
)

func TestCreateUserQuery(t *testing.T) {
	testCases := []struct {
		name   model.UserName
	}{
		{name: "hoge"},
	}

	for _, tc := range testCases {
		repository := &UserRepositoryImpl{}
		got,err := repository.CreateUserQuery(tc.name)
		want := tc.name

		if err != nil {
			t.Fatalf("CreateUserQuery() error = %v", err)
		}
		if got.UserName != tc.name{
			t.Errorf("postgresql.CreateUserQuery() = %v, want %v",got.UserName,want)
		}
}}

func TestUpdateUserQuery(t *testing.T) {
	testCases := []struct{
		initialName model.UserName
		updateName model.UserName
	}{
		{initialName: "hoge",updateName: "hogehoge"},
	}
	for _, tc := range testCases {
		repository := &UserRepositoryImpl{}
		user,err := repository.CreateUserQuery(tc.initialName)
		if err != nil {
			t.Fatalf("CreateUserQuery() error = %v", err)
		}
		
		updateUser := model.User{
			UserId: user.UserId,
			UserName: tc.updateName,
		}
		got,err := repository.UpdateUserQuery(updateUser)
		want := updateUser

		if err != nil {
			t.Fatalf("CreateUserQuery() error = %v", err)
		}
		if got.UserName != tc.updateName{
			t.Errorf("postgresql.UpdateUserQuery() = %v, want %v", got, want)
		}
}}
