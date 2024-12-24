package model

type RoomId int

type RoomName string

type CreateRoom struct {
	IsHuman  bool
	RoomName RoomName
	UserId   UserId
}
