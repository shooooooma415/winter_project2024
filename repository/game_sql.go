package repository

import (
	"winter_pj/model"
)

type GameRepository interface {
	GetCompareImage(searchName string) (*model.ImageList, error)
	InsertCompareImage(insertCompareImage model.InsertCompareImage) error
	InsertFaceImage(userId model.UserId, image string) error
}
