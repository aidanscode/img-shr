package model

import "time"

type Post struct {
	Id int
	AuthorId string
	Title string
	ImgPath string
	CreatedAt time.Time
}
