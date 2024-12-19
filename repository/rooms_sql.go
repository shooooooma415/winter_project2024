package repository

import (
	"winter_pj/model"
)

type RoomsRepository interface {
	CreateRoom(isHuman bool, password model.PassWord, userId model.UserId) (model.RoomId, error)
	JoinRoom(password model.PassWord, userId model.UserId) (model.RoomId, error)
	GetRoomId(roomId model.RoomId) (model.PassWord, error)
	DeleteRoom(roomId model.RoomId)(bool,error)
}
