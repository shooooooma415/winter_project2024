package model

type UserId int

//ユーザー作成のリクエスト
type CreateUserRequest struct{
	Name string `json:"name"`
}

//ユーザー作成のレスポンス
type CreateUserResponse struct{
	UserId UserId `json:"user_id"`
}

//ユーザー名変更のリクエスト，レスポンスはAuthResponse使う
type UpdateUserRequest struct{
	Name string `json:"name"`
}