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
	testDateCommentListByPostId = time.Date(2018, time.September, 16, 12, 0, 0, 0, time.UTC)
)

var validCommentsCommentListByPostId = []model.Comment{
	{Id: 123, PostId: 3, Comment: "abc", Author: "cool author", CreationDate: testDateCommentListByPostId},
	{Id: 321, PostId: 3, Comment: "def", Author: "cool author2", CreationDate: testDateCommentListByPostId},
	{Id: 543, PostId: 3, Comment: "ghi", Author: "cool author3", CreationDate: testDateCommentListByPostId},
}

func TestCommentListByPostIdRest(t *testing.T) {
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
			testName:           "testSuccessfullyGetComments",
			commentRepository:  repository.CustomCommentRepository(validCommentsCommentListByPostId),
			postRepository:     repository.CustomPostRepository(make([]model.Post, 0)),
			postId: 3,
			expectedHttpStatus: 200,
			expectedHeader:     "application/json",
			expectedResponse:   validCommentsCommentListByPostId,
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(tt *testing.T) {
			// GIVEN
			svc := RestApiService{commentRepository: &tc.commentRepository,
				postRepository: &tc.postRepository}

			req := httptest.NewRequest(http.MethodGet, "https://example.com/api/get/comments/"+strconv.Itoa(tc.postId), nil)
			w := httptest.NewRecorder()

			// WHEN
			handleCommentListByPostId(&svc)(w, req)
			response := w.Result()
			body, _ := io.ReadAll(response.Body)
			var commentsList []model.Comment
			err := json.Unmarshal(body, &commentsList)

			if err != nil {
				t.Fail()
			}

			// THEN
			assert.Equal(t, tc.expectedHttpStatus, response.StatusCode)
			assert.Equal(t, tc.expectedHeader, response.Header.Get("Content-Type"))
			assert.ElementsMatch(t, tc.expectedResponse, commentsList)
		})
	}
}