package model

type ImageId int

type ImageName string

type Image string //base64で画像をやり取りするから文字列

type CompareImage string //これってImageと分ける必要あるんですかね？？

type CompareImageName int //これも分けないでいいかも

type Ranking struct{
	Position int
	UserId UserId
	UserName UserName
	FaceImage Image
}