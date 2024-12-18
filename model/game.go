package model

type ImageId int

type Ranking struct{
	Position int
	UserId UserId
	UserName string
	FaceImage string
}