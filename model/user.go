package model

//ユーザー作成のリクエスト
type CreateUserRequest struct{
	Name string `json:"name"`
}

//ユーザー作成のレスポンス
type CreateUserResponse struct{
	UserId int `json:"user_id"`
}

//ユーザー名変更のリクエスト，レスポンスはAuthResponse使う
type RenewUserRequest struct{
	Name string `json:"name"`
}