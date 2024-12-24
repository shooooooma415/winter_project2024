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
