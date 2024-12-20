package game

import (
	"winter_pj/model"
)

type GameRepository interface {
	GetCompareImage(searchName string) (*model.ImageList, error)
	InsertCompareImage(insertCompareImage model.InsertCompareImage) error
	InsertFaceImage(insertFaceImage model.InsertFaceImage) error
}
