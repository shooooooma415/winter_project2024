package room

import (
	"winter_pj/model"
)

type RoomsRepository interface {
	CreateRoom(createRoom model.CreateRoom) (*model.Room, error)
	JoinRoom(roomId model.RoomId, userId model.UserId) (*model.JoinRoom, error)
	GetRoomId(roomId model.RoomId) (*model.RoomName, error)
	DeleteRoom(roomId model.RoomId) error
}
