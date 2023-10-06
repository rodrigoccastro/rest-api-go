package service

import (
	"bytes"
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

func TestCommentAddRest(t *testing.T) {
	tests := []struct {
		testName           string
		commentRepository  repository.CommentRepository
		postRepository     repository.PostRepository
		comment            model.Comment
		expectedHttpStatus int
		expectedHeader     string
		expectedResponse   interface{}
	}{
		{
			testName:           "testSuccessfullyAddComment",
			commentRepository:  repository.CustomCommentRepository(make([]model.Comment, 0)),
			postRepository:     repository.CustomPostRepository(make([]model.Post, 0)),
			comment:            model.Comment{Id: 123, PostId: 3, Comment: "cool cmnt", Author: "cool auth", CreationDate: time.Now()},
			expectedHttpStatus: 200,
			expectedHeader:     "application/json",
			expectedResponse:   RestJsonResponse{Message: "comment id: 123 successfully added", Status: http.StatusOK},
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(tt *testing.T) {
			// GIVEN
			svc := RestApiService{commentRepository: &tc.commentRepository,
				postRepository: &tc.postRepository}

			data, _ := json.Marshal(&tc.comment)
			req := httptest.NewRequest(http.MethodPost, "https://example.com/api/post/comment", bytes.NewReader(data))
			w := httptest.NewRecorder()

			// WHEN
			handleCommentAdd(&svc)(w, req)

			response := w.Result()
			body, _ := io.ReadAll(response.Body)
			var resp RestJsonResponse
			err := json.Unmarshal(body, &resp)

			if err != nil {
				t.Fail()
			}

			// THEN
			assert.Equal(t, tc.expectedHttpStatus, response.StatusCode)
			assert.Equal(t, tc.expectedHeader, response.Header.Get("Content-Type"))
			assert.Equal(t, tc.expectedResponse, resp)
		})
	}
}