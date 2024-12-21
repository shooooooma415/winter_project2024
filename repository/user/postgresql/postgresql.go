package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"winter_pj/model"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

func (q *UserRepositoryImpl) CreateUserQuery(name model.UserName) (*model.User, error) {
	query := `
		INSERT INTO users (name) 
		VALUES ($1)
		RETURNING id, name
	`

	var user model.User
	err := q.DB.QueryRow(query, name).Scan(&user.UserId, &user.UserName)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &user, nil
}

func (q *UserRepositoryImpl) UpdateUserQuery(userID model.UserId, name model.UserName) (*model.User, error) {
	query := `
		UPDATE users
		SET name = $2
		WHERE id = $1
		RETURNING name, id
	`

	var updatedUser model.User
	err := q.DB.QueryRow(query, userID, name).Scan(&updatedUser.UserId, &updatedUser.UserName)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("no rows were updated: user with ID %d not found", userID)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	return &updatedUser, nil
}
