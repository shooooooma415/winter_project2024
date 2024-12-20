package user

import (
	"winter_pj/model"
)

type UsersRepository interface {
	CreateUser(name model.UserName) (*model.UserId, error)
	UpdateUser(userId model.UserId, name model.UserName) (*model.User, error)
}