package postgresql

import (
	"database/sql"
	"fmt"
	"winter_pj/model"
)

type RoomRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *RoomRepositoryImpl {
	return &RoomRepositoryImpl{DB: db}
}

func (q *RoomRepositoryImpl) CreateRoomQuery(createRoom model.CreateRoom) (*model.Room, error) {
	query := `
	INSERT INTO rooms (author_id,name)
	VALUES ($2,$1)
	RETURNING id,author_id,name
	`

	var room model.Room
	err := q.DB.QueryRow(
		query,
		createRoom.AuthorId,
		createRoom.RoomName,
	).Scan(
		&room.RoomId,
		&room.AuthorId,
		&room.RoomName,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create room: %w", err)
	}
	return &room, nil
}

func (q *RoomRepositoryImpl) JoinRoomQuery(roomMember model.RoomMember) (*model.RoomMember, error) {
	query := `
		INSERT INTO room_member (user_id)
		VALUES $2
		WHERE room_id = $1
		RETURNING room_id,user_id
	`

	var room model.RoomMember
	err := q.DB.QueryRow(query, roomMember.RoomId, roomMember.UserId).Scan(&room.RoomId, &room.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to join room: %w", err)
	}
	return &room, nil
}

func (q *RoomRepositoryImpl) GetAuthorIdQuery(roomId model.RoomId) (*model.UserId, *model.RoomId, error) {
	query := `
	SELECT author_id, id
	WHERE id = $1
	`

	var authorId model.UserId
	var returnRoomId model.RoomId
	err := q.DB.QueryRow(query, roomId).Scan(authorId, returnRoomId)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to delete room: %w", err)
	}
	return &authorId,&roomId,nil
}

func (q *RoomRepositoryImpl) DeleteRoomQuery(roomId model.RoomId) (*model.RoomId, error) {
	query := `
		DELETE FROM rooms
		WHERE room_id = $1
	`

	var returnRoomId model.RoomId
	err := q.DB.QueryRow(query, roomId).Scan(returnRoomId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete room: %w", err)
	}
	return &returnRoomId, nil
}

func (q *RoomRepositoryImpl) WithdrawRoomQuery(roomMember model.RoomMember) (*model.RoomMember, error) {
	query := `
		DELETE FROM room_member
		WHERE room_id = $1
		AND user_id = $2
		RETURNING room_id
	`

	var room model.RoomMember
	err := q.DB.QueryRow(query, roomMember.RoomId, roomMember.UserId).Scan(&room.RoomId, &room.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to withdraw room: %w", err)
	}
	return &room, nil
}
