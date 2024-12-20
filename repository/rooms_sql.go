package repository

import (
	"winter_pj/model"
)

type RoomsRepository interface {
	CreateRoom(createRoom model.CreateRoom) (*model.RoomId, error)
	JoinRoom(joinRoom model.JoinRoom) (*model.RoomId, error)
	GetRoomId(roomId model.RoomId) (*model.Password, error)
	DeleteRoom(roomId model.RoomId)(error)
}
