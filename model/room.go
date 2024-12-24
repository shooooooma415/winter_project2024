package model

type RoomId int

type RoomName string

type CreateRoom struct {
	IsHuman  bool
	RoomName RoomName
	AuthorId UserId
}

type Room struct {
	IsHuman  bool
	RoomName RoomName
	AuthorId UserId
	RoomId   RoomId
}
