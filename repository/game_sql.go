package repository

import (
	"winter_pj/model"
)

// 実行結果boolで表すんならerrorいらない？いいやり方がある？
type GameRepository interface {
	GetCompareImage(searchName string) (model.ImageList, error)
	InsertCompareImage(imageName string, image string, isHuman bool) (error)
	InsertFaceImage(userId model.UserId, image string) (error)
}
