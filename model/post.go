package model

import (
	"fmt"
	"time"
)

type Post struct {
	Id int
	Title string
	ImgPath string
	CreatedAt time.Time
}

func NewPost(title string) *Post {
	p := &Post{Title: title, CreatedAt: time.Now()}
	return p
}

func (p *Post) GetPostUrl() string {
	return fmt.Sprintf("/i/%d", p.Id)
}

func (p *Post) GetDirectUrl() string {
	return fmt.Sprintf("/i/%d/raw", p.Id)
}
