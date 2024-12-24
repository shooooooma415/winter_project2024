package room

import (
	"winter_pj/model"
)

type RoomsRepository interface {
	CreateRoom(createRoom model.CreateRoom) (*model.Room, error)
	JoinRoom(password model.RoomName, userId model.UserId) (*model.RoomId, error)
	GetRoomId(roomId model.RoomId) (*model.RoomName, error)
	DeleteRoom(roomId model.RoomId) error
}
