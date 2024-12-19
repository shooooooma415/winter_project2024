package model

type ImageId int

type Ranking struct {
	Position  int
	UserId    UserId
	UserName  string
	FaceImage string
}

type ImageList struct {
	ImageId   ImageId `json:"image_id"`
	ImageName string  `json:"image_name"`
	Image     string  `json:"image"`
}
