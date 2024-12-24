package postgresql

import (
	"database/sql"
	"winter_pj/model"
)

type RoomRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *RoomRepositoryImpl {
	return &RoomRepositoryImpl{DB: db}
}

func (q *RoomRepositoryImpl) CreateRoomQuery(createRoom model.CreateRoom) (model.RoomId, error) {
	
}
