package repository

import (
	"testing"
	"time"

	"github.com/rodrigoccastro/rest-api-go/model"
	"github.com/stretchr/testify/assert"
)

var (
	post1          = model.Post{Id: 1, Title: "title 1", Content: "content 1", CreationDate: time.Unix(10011, 0)}
	post2          = model.Post{Id: 2, Title: "title 2", Content: "content 2", CreationDate: time.Unix(10011, 0)}
)

func TestPostInsert(t *testing.T) {
	c := PostRepository{}
	if err := c.Insert(post1); err != nil {
		t.Fail()
	}
	err := c.Insert(post1)
	assert.EqualErrorf(t, err, "Error: Post with id: 1 already exists in the repository!", "test failed because of wrong error msg: %+v", err)
}

func TestPostGetById(t *testing.T) {
	c := PostRepository{}
	
	_, err := c.GetById(post1.Id)
	if err==nil {
		t.Fail()
	}
	
	c.Insert(post1);
	post, _ := c.GetById(post1.Id)
	if post==nil {
		t.Fail()
	}

	// change dates for test
	post1.CreationDate = post.CreationDate
	assert.Equal(t, &post1, post)
}

func TestPostGetAll(t *testing.T) {
	c := PostRepository{}
	c.Insert(post1)
	c.Insert(post2)

	result := c.GetAll()
	post1.CreationDate = result[0].CreationDate
	post2.CreationDate = result[1].CreationDate

	assert.Equal(t, post1, result[0])
	assert.Equal(t, post2, result[1])
}


func TestUpdate(t *testing.T) {
	c := PostRepository{}
	c.Insert(post1);

	post1.Title = "another title"
	post1.Content = "another content"

	err := c.Update(post1)
	if err!=nil {
		t.Fail()
	}

	obj, _ := c.GetById(post1.Id)
	assert.Equal(t, obj.Title, "another title")
	assert.Equal(t, obj.Content, "another content")
}
