package model

import "time"

type Post struct {
	Id int
	AuthorId int
	Title string
	ImgPath string
	CreatedAt time.Time
}

func NewPost(authorId int, title string, imgPath string) *Post {
	p := &Post{AuthorId: authorId, Title: title, ImgPath: imgPath, CreatedAt: time.Now()}
	return p
}
