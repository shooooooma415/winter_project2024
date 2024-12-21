package user

import (
	"winter_pj/model"
)

type UsersRepository interface {
	CreateUser(name model.UserName) (*model.User, error)
	UpdateUser(user model.User) (*model.User, error)
}