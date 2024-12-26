package room

import (
	"winter_pj/model"
)

type RoomRepository interface {
	CreateRoom(createRoom model.CreateRoom) (*model.Room, error)
	JoinRoom(roomMember model.RoomMember) (*model.RoomMember, error)
	GetRoomName(roomId model.RoomId) (*model.RoomName, *model.RoomId, error)
	GetAuthorId(roomId model.RoomId) (*model.UserId, *model.RoomId, error)
	DeleteRoom(roomId model.RoomId) (*model.RoomId, error)
	DeleteRoomMemberTable(roomId model.RoomId)(*model.RoomId, error)
	WithdrawRoom(roomMember model.RoomMember) (*model.RoomMember, error)
}
