package model

type ImageId int

// リストの要素
type CompareImage struct {
	ImageID   ImageId `json:"image_id"`
	ImageName string  `json:"image_name"`
	Image     string  `json:"image"` // base64形式の文字列がはいる
}

// 比較画像の選択のリクエスト．レスポンスはAuthResponseを使う
type SelectImageRequest struct {
	ImageId ImageId `json:"image_id"`
}

// 比較画像を渡す際のレスポンス
type GetCompareImageRequest struct {
	Images []CompareImage `json:"images"`
}

// 比較画像の挿入のレスポンス
type InsertCompareImageRequest struct {
	ImageName string `json:"image_name"`
	Image     string `json:"image"` // base64形式の文字列がはいる
}

// 比較画像の挿入のレスポンス
type InsertCompareImageResponse struct {
	ImageId ImageId `json:"image_id"`
}

// 撮影した画像の挿入のリクエスト．レスポンスはAuthResponseを使う
type InsertFaceImageRequest struct {
	Image string `json:"image"` // base64形式の文字列がはいる
}

// ランキングの要素の構造体
type Ranking struct {
	Position  int    `json:"position"`
	UserID    UserId `json:"user_id"`
	UserName  string `json:"user_name"`
	FaceImage string `json:"face_image"`
}

// 結果のレスポンス
type ResultRankingResponse struct {
	CompareImage     string    `json:"compare_image"`
	CompareImageName string    `json:"compare_image_name"`
	Ranking          []Ranking `json:"ranking"`
}
