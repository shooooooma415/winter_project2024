package model

type RoomId int

//ルーム作成のリクエスト
type CreateRoomRequest struct{
	Password string `json:"password"` 
	UserId UserId `json:"user_id"`
}

//ルーム作成のレスポンス
type CreateRoomResponse struct{
	RoomId RoomId `json:"room_id"`
	IsAuthor bool `json:"is_author"`
}

//ルーム参加のリクエスト
type JoinRoomRequest struct{
	Password string `json:"password"`
	UserId UserId `json:"user_id"`
}

//ルーム参加のレスポンス
type JoinRoomResponse struct{
	RoomId RoomId `json:"room_id"`
	IsAuthor bool `json:"is_author"`
}