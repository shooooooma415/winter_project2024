package model

//ルーム作成のリクエスト
type CreateRoomRequest struct{
	Password string `json:"password"` 
	UserId int `json:"user_id"`
}

//ルーム作成のレスポンス
type CreateRoomResponse struct{
	RoomId int `json:"room_id"`
	IsAuthor bool `json:"is_author"`
}

//ルーム参加のリクエスト
type JoinRoomRequest struct{
	Password string `json:"password"`
	UserId int `json:"user_id"`
}

//ルーム参加のレスポンス
type JoinRoomResponse struct{
	RoomId int `json:"room_id"`
	IsAuthor bool `json:"is_author"`
}