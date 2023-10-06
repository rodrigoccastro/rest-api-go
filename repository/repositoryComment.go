package repository

import (
	"fmt"

	"github.com/rodrigoccastro/rest-api-go/model"
)

type CommentRepository struct {
	repository []model.Comment
}

func NewCommentRepository() *CommentRepository {
	repo := CustomCommentRepository(make([]model.Comment, 0))
	return &repo
}

func CustomCommentRepository(mockStorage []model.Comment) CommentRepository {
	return CommentRepository{repository: mockStorage}
}

type CommentAlreadyExistsError struct {
	id uint64
}

func (e CommentAlreadyExistsError) Error() string {
	return fmt.Sprintf("Error: Comment with id: %v already exists in the repository!", e.id)
}

type CommentNotFoundError struct {
	id uint64
}

func (e CommentNotFoundError) Error() string {
	return fmt.Sprintf("Error: Comment with id: %v was not found in the repository!", e.id)
}

func (c *CommentRepository) Insert(comment model.Comment) error {
	v, _ := c.GetById(comment.Id)
	if (v != nil) {
		return CommentAlreadyExistsError {id: comment.Id}
	}
	c.repository = append(c.repository, comment)
	return nil
}

func (c *CommentRepository) GetById(id uint64) (*model.Comment, error) {
	for i := 0; i < len(c.repository); i++ {
		item := &c.repository[i]
		if (item.Id == id) {
			return item, nil
		}
	}
	return nil, CommentNotFoundError {id: id}
}

func (c *CommentRepository) GetAllByPostId(id uint64) []model.Comment {
	ret := []model.Comment{}
	for i := 0; i < len(c.repository); i++ {
		item := c.repository[i]
		if (item.PostId == id) {
			ret = append(ret, item)
		}
	}
	return ret
}