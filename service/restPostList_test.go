package service

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rodrigoccastro/rest-api-go/model"
	"github.com/rodrigoccastro/rest-api-go/repository"
	"github.com/stretchr/testify/assert"
)

var testDateListRest = time.Date(2018, time.September, 16, 12, 0, 0, 0, time.UTC)
var validPostListRest1 = model.Post{Id: 1, Title: "happy post1", Content: "test content1", CreationDate: testDateListRest}
var validPostListRest2 = model.Post{Id: 2, Title: "happy post2", Content: "test content2", CreationDate: testDateListRest}

func TestPostListRest(t *testing.T) {
	t.Run("testSuccessfullyListPost", func(tt *testing.T) {
		// GIVEN
		svc := RestApiService{commentRepository: nil,
				postRepository: repository.NewPostRepository()}

		req := httptest.NewRequest(http.MethodGet, "https://example.com/api/post/list", nil)
		w := httptest.NewRecorder()

		// WHEN
		svc.postRepository.Insert(validPostListRest1)
		svc.postRepository.Insert(validPostListRest2)
		handlePostList(&svc)(w, req)

		response := w.Result()
		body, _ := io.ReadAll(response.Body)
		var postsList []model.Post
		err := json.Unmarshal(body, &postsList)

		if err != nil {
			t.Fail()
		}

		// THEN
		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, "application/json", response.Header.Get("Content-Type"))

		assert.Equal(t, validPostListRest1.Id, postsList[0].Id)
		assert.Equal(t, validPostListRest2.Id, postsList[1].Id)
	})
}
