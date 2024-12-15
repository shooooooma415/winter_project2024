package repository

import (
	"model"
	"fmt"
)

type userRepo struct{}

type UserRepo interface{
	
}

func NewUserRepo() UserRepo{
	return userRepo{}
}