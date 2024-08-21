package service

import "github.com/AidansCode/img-shr/model"

type PostService interface {
	Latest(n int) []model.Post
}

type postService struct {
}

func NewPostService() PostService {
	return &postService{}
}

func (ps *postService) Latest(n int) []model.Post {
	posts := []model.Post{
		{Id: 1, AuthorId: "Aidan", Title: "By Me", ImgPath: "https://placehold.co/300x200"},
		{Id: 2, AuthorId: "Ricky", Title: "Ricky Original", ImgPath: "https://placehold.co/300x200"},
	}
	return posts
}
