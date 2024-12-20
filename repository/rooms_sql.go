package repository

import (
	"winter_pj/model"
)

type RoomsRepository interface {
	CreateRoom(isHuman bool, password model.Password, userId model.UserId) (*model.RoomId, error)
	JoinRoom(password model.Password, userId model.UserId) (*model.RoomId, error)
	GetRoomId(roomId model.RoomId) (*model.Password, error)
	DeleteRoom(roomId model.RoomId)(error)
}
