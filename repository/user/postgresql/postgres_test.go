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
		if err != nil{
			t.Errorf("Error")
		}
	}
}

func TestUpdateUserQuery(t *testing.T) {}
