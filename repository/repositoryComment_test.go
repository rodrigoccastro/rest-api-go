package repository

import (
	"testing"
	"time"

	"github.com/rodrigoccastro/rest-api-go/model"
	"github.com/stretchr/testify/assert"
)

var (
	comment1          = model.Comment{Id: 1, PostId: 101, Comment: "comment2", Author: "author2", CreationDate: time.Unix(10011, 0)}
	comment2          = model.Comment{Id: 2, PostId: 101, Comment: "comment2", Author: "author2", CreationDate: time.Unix(10011, 0)}
	comment3          = model.Comment{Id: 3, PostId: 100, Comment: "comment3", Author: "author3", CreationDate: time.Unix(10022, 0)}
	NonExistentPostId = uint64(10101010)
)

func TestCommentInsert(t *testing.T) {
	c := CommentRepository{}
	if err := c.Insert(comment1); err != nil {
		t.Fail()
	}
	err := c.Insert(comment1)
	assert.EqualErrorf(t, err, "Error: Comment with id: 1 already exists in the repository!", "test failed because of wrong error msg: %+v", err)
}

func TestCommentGetById(t *testing.T) {
	c := CommentRepository{}
	
	_, err := c.GetById(comment1.Id)
	if err==nil {
		t.Fail()
	}
	
	c.Insert(comment1);
	comment, _ := c.GetById(comment1.Id)
	if comment==nil {
		t.Fail()
	}

	// change dates for test
	comment1.CreationDate = comment.CreationDate
	assert.Equal(t, &comment1, comment)
}

func TestCommentGetAllByPostId(t *testing.T) {
	c := CommentRepository{}
	c.Insert(comment1)
	c.Insert(comment2)
	c.Insert(comment3)

	expectedResult := []model.Comment{comment1, comment2}
	result := c.GetAllByPostId(comment1.PostId)
	assert.ElementsMatch(t, expectedResult, result)
}