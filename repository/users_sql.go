package repository

import (
	"database/sql"
	"fmt"
	"winter_pj/model"
)

type UsersRepository interface {
	CreateUser(name string) (*model.UserId, error)
	UpdateUser(userId model.UserId, name model.UserName) (*model.User, error)
}

type UsersSQL struct {
	DB *sql.DB
}

func NewUsersSQL(db *sql.DB) *UsersSQL {
	return &UsersSQL{DB: db}
}

func (q *UsersSQL) CreateUserQuery(name string) (*model.UserId, error) {
	query := `
		INSERT INTO users (name) 
		VALUES ($1)
		RETURNING id
	`

	var userID model.UserId
	err := q.DB.QueryRow(query, name).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &userID, nil
}

func (q *UsersSQL) UpdateUserQuery(userID model.UserId, name model.UserName) (*model.User,error) {
	query := `
		UPDATE users
		SET name = $2
		WHERE id = $1
	`

	result, err := q.DB.Exec(query, userID, name)
	if err != nil {
		return nil,fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil,fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return nil,fmt.Errorf("no rows were updated: user with ID %d not found", userID)
	}
	return &model.User{
		UserId:   userID,
		UserName: name,
	}, nil
}
