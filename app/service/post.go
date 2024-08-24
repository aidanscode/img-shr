package service

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/AidansCode/img-shr/app/db"
	"github.com/AidansCode/img-shr/model"
)

var (
	ErrNotFound = errors.New("post not found with given id")
)

type PostService interface {
	Latest(n uint) ([]model.Post, error)
	Get(id int) (*model.Post, error)
	Save(p *model.Post, file *multipart.File, validatedMimeType string) (*model.Post, error)
	Update(p *model.Post) (*model.Post, error)
}

type postService struct {
	Database db.Database
	newImageDirectory string
}

func NewPostService(database *db.Database, newImageDirectory string) PostService {
	return &postService{Database: *database, newImageDirectory: newImageDirectory}
}

func (ps *postService) Latest(n uint) ([]model.Post, error) {
	rows, err := ps.Database.Db.Query("SELECT * FROM posts WHERE img_path IS NOT NULL ORDER BY created_at DESC LIMIT ?", n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.Id, &post.Title, &post.ImgPath, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (ps *postService) Get(id int) (*model.Post, error) {
	rows, err := ps.Database.Db.Query("SELECT * FROM posts WHERE img_path IS NOT NULL AND id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, ErrNotFound
	}

	var post model.Post
	err = rows.Scan(&post.Id, &post.Title, &post.ImgPath, &post.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (ps *postService) Save(p *model.Post, file *multipart.File, validatedMimeType string) (*model.Post, error) {
	p, err := ps.saveToDb(p)
	if err != nil {
		return p, err
	}

	filePath, err := ps.saveFile(p.Id, file, validatedMimeType)
	if err != nil {
		return p, err
	}

	p.ImgPath = filePath
	p, err = ps.Update(p)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (ps *postService) Update(p *model.Post) (*model.Post, error) {
	if p.Id == 0 {
		return p, errors.New("post hasn't been saved yet")
	}

	_, err := ps.Database.Db.Exec("UPDATE posts SET title=?, img_path=?, created_at=? WHERE id=?", p.Title, p.ImgPath, p.CreatedAt, p.Id)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (ps *postService) saveToDb(p *model.Post) (*model.Post, error) {
	if p.Id != 0 {
		return p, errors.New("post has already been saved")
	}

	res, err := ps.Database.Db.Exec("INSERT INTO posts VALUES (null, ?, NULL, ?)", p.Title, p.CreatedAt)
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

func (ps *postService) saveFile(id int, uploadedFile *multipart.File, validatedMimeType string) (string, error) {
	fileExtension := strings.Split(validatedMimeType, "/")[1]

	path, err := filepath.Abs(fmt.Sprintf("%s/%d.%s", ps.newImageDirectory, id, fileExtension))
	if err != nil {
		return "", err
	}

	newFile, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer newFile.Close()

	if _, err := io.Copy(newFile, *uploadedFile); err != nil {
		return "", err
	}

	return path, nil
}
