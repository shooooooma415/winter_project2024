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
		if got.UserName != tc.name{
			t.Errorf("postgresql.CreateUserQuery() = %v, want %v",got.UserName,tc.name)
		}
}

func TestUpdateUserQuery(t *testing.T) {
	testCases := []{
		use model.User
	}
}
