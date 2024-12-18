package repository

//modelのPR通ったら型を変更する
type UsersRepository interface{
	CreateUser(name string) (int, error)
	UpdateUser(user_id int, name string)
}

type UsersSQL struct{}

func NewUsersSQL() *UsersSQL {
	return &UsersSQL{}
}

// これ雛形
// func (q *UsersSQL) GetAllUsersQuery() string {
// 	return `
// 		SELECT id, name, email
// 		FROM users
// 	`
// }

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
