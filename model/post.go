package model

import (
	"fmt"
	"time"
)

type Post struct {
	Id int
	AuthorId int
	Title string
	ImgPath string
	CreatedAt time.Time
}

func NewPost(authorId int, title string) *Post {
	p := &Post{AuthorId: authorId, Title: title, CreatedAt: time.Now()}
	return p
}

func (p *Post) GetDirectUrl() string {
	return fmt.Sprintf("/i/%d/raw", p.Id)
}
