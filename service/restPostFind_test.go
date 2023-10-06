package service

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/rodrigoccastro/rest-api-go/model"
	"github.com/rodrigoccastro/rest-api-go/repository"
	"github.com/stretchr/testify/assert"
)

var (
	testDateFindRest = time.Date(2018, time.September, 16, 12, 0, 0, 0, time.UTC)
)

var validPostFindRest = model.Post{Id: 34, Title: "happy post", Content: "test content", CreationDate: testDateFindRest}

func TestPostFindRest(t *testing.T) {
	tests := []struct {
		testName           string
		commentRepository  repository.CommentRepository
		postRepository     repository.PostRepository
		postId             int
		expectedHttpStatus int
		expectedHeader     string
		expectedResponse   interface{}
	}{
		{
			testName:           "testSuccessfullyGetPost",
			commentRepository:  repository.CustomCommentRepository(make([]model.Comment, 0)),
			postRepository:     repository.CustomPostRepository([]model.Post{validPostFindRest}),
			postId: int(validPostFindRest.Id),
			expectedHttpStatus: 200,
			expectedHeader:     "application/json",
			expectedResponse:   validPostFindRest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(tt *testing.T) {
			// GIVEN
			svc := RestApiService{commentRepository: &tc.commentRepository,
				postRepository: &tc.postRepository}

			req := httptest.NewRequest(http.MethodGet, "https://example.com/api/get/post/"+strconv.Itoa(tc.postId), nil)
			w := httptest.NewRecorder()

			// WHEN
			handlePostFind(&svc)(w, req)
			response := w.Result()
			body, _ := io.ReadAll(response.Body)
			var post model.Post
			err := json.Unmarshal(body, &post)

			if err != nil {
				t.Fail()
			}

			// THEN
			assert.Equal(t, tc.expectedHttpStatus, response.StatusCode)
			assert.Equal(t, tc.expectedHeader, response.Header.Get("Content-Type"))
			assert.Equal(t, tc.expectedResponse, post)
		})
	}
}
