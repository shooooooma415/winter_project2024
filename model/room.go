package model

type RoomId int

type Password string

type CreateRoom struct {
	IsHuman  bool
	Password Password
	UserId   UserId
}

type JoinRoom struct {
	Password Password
	UserId   UserId
}
