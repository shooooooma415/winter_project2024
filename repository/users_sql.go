package repository

import (
	"database/sql"
	"fmt"
	"winter_pj/model"
)

type UsersRepository interface {
	CreateUser(name string) (*model.UserId, error)
	UpdateUser(userId model.UserId, name string) (error)
}

type UsersSQL struct {
	DB *sql.DB
}

func NewUsersSQL(db *sql.DB) *UsersSQL {
	return &UsersSQL{DB: db}
}

// 調べた書き方なのであってるかはわからない
func (q *UsersSQL) CreateUserQuery(name string) (model.UserId, error) {
	query := `
		INSERT INTO users (name) 
		VALUES ($1)
		RETURNING id
	`

	var userID model.UserId
	err := q.DB.QueryRow(query, name).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}
	return userID, nil
}

func (q *UsersSQL) UpdateUserQuery(userID model.UserId, name string) (bool, error) {
	query := `
		UPDATE users
		SET name = $2
		WHERE id = $1
	`

	result, err := q.DB.Exec(query, userID, name)
	if err != nil {
		return false, fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("failed to check rows affected: %w", err)
	}

	return rowsAffected > 0, nil
}

// // クエリ: 特定のユーザーを取得
// func (q *UsersSQL) GetUserByIDQuery() string {
// 	return `
// 		SELECT id, name, email
// 		FROM users
// 		WHERE id = $1
// 	`
// }

// // クエリ: ユーザーを追加
// func (q *UsersSQL) CreateUserQuery() string {
// 	return `
// 		INSERT INTO users (name, email)
// 		VALUES ($1, $2)
// 		RETURNING id
// 	`
// }
