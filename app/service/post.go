package service

import (
	"errors"

	"github.com/AidansCode/img-shr/app/db"
	"github.com/AidansCode/img-shr/model"
)

type PostService interface {
	Latest(n int) []model.Post
	Save(p *model.Post) (*model.Post, error)
}

type postService struct {
	Database db.Database
}

func NewPostService(database *db.Database) PostService {
	return &postService{Database: *database}
}

func (ps *postService) Latest(n int) []model.Post {
	posts := []model.Post{
		{Id: 1, AuthorId: 1, Title: "By Me", ImgPath: "https://placehold.co/300x200"},
		{Id: 2, AuthorId: 2, Title: "Ricky Original", ImgPath: "https://placehold.co/300x200"},
	}
	return posts
}

func (ps *postService) Save(p *model.Post) (*model.Post, error) {
	if p.Id != 0 {
		return p, errors.New("post has already been saved")
	}

	res, err := ps.Database.Db.Exec("INSERT INTO posts VALUES (null, ?, ?, ?, ?)", p.AuthorId, p.Title, p.ImgPath, p.CreatedAt)
	if err != nil {
		return p, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return p, err
	}

	p.Id = int(id)

	return p, nil
}
