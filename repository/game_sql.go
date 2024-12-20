package repository

import (
	"winter_pj/model"
)

type GameRepository interface {
	GetCompareImage(searchName string) (*model.ImageList, error)
	InsertCompareImage(imageName string, image string, isHuman bool) (error)
	InsertFaceImage(userId model.UserId, image string) (error)
}
