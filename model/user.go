package model

type CreateUserRequest struct{
	Name string `json:"name"`
}

type CreateUserResponse struct{
	UserId int `json:"user_id"`
}

type RenewUserRequest struct{
	Name string `json:"name"`
}