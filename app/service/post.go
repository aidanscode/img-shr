package service

import (
	"errors"

	"github.com/AidansCode/img-shr/app/db"
	"github.com/AidansCode/img-shr/model"
)

type PostService interface {
	Latest(n uint) ([]model.Post, error)
	Save(p *model.Post) (*model.Post, error)
}

type postService struct {
	Database db.Database
	newImageDirectory string
}

func NewPostService(database *db.Database, newImageDirectory string) PostService {
	return &postService{Database: *database, newImageDirectory: newImageDirectory}
}

func (ps *postService) Latest(n uint) ([]model.Post, error) {
	rows, err := ps.Database.Db.Query("SELECT * FROM posts ORDER BY created_at DESC LIMIT ?", n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.Id, &post.AuthorId, &post.Title, &post.ImgPath, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
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
