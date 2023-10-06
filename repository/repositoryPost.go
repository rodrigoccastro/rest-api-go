package repository

import (
	"fmt"
	"time"

	"github.com/rodrigoccastro/rest-api-go/model"
)

type PostRepository struct {
	repository []model.Post
}

func CustomPostRepository(mockStorage []model.Post) PostRepository {
	return PostRepository{repository: mockStorage}
}

func NewPostRepository() *PostRepository {
	repo := CustomPostRepository(make([]model.Post, 0))
	return &repo
}

type PostAlreadyExistsError struct {
	id uint64
}

func (e PostAlreadyExistsError) Error() string {
	return fmt.Sprintf("Error: Post with id: %v already exists in the repository!", e.id)
}

type PostNotFoundError struct {
	id uint64
}

func (e PostNotFoundError) Error() string {
	return fmt.Sprintf("Error: Post with id: %v was not found in the repository!", e.id)
}

func (c *PostRepository) Insert(post model.Post) error {
	v, _ := c.GetById(post.Id)
	if (v != nil) {
		return PostAlreadyExistsError {id: post.Id}
	}
	post.CreationDate = time.Now()
	c.repository = append(c.repository, post)
	return nil
}

func (c *PostRepository) GetById(id uint64) (*model.Post, error) {
	for i := 0; i < len(c.repository); i++ {
		item := &c.repository[i]
		if (item.Id == id) {
			return item, nil
		}
	}
	return nil, PostNotFoundError {id: id}
}

func (c *PostRepository) GetAll() []model.Post {
	return c.repository
}

func (c *PostRepository) Update(post model.Post) error {
	for i := 0; i < len(c.repository); i++ {
		item := &c.repository[i]
		if (item.Id == post.Id) {
			item.Id = post.Id 
			item.Title = post.Title
			item.Content = post.Content
			return nil
		}
	}
	return PostAlreadyExistsError {id: post.Id}
}