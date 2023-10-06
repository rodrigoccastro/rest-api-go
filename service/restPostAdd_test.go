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

func TestPostAddRest(t *testing.T) {
	tests := []struct {
		testName           string
		commentRepository  repository.CommentRepository
		postRepository     repository.PostRepository
		post               interface{}
		expectedHttpStatus int
		expectedHeader     string
		expectedResponse   interface{}
	}{
		{
			testName:           "testSuccessfullyAddPost",
			post:               model.Post{Id: 256, Title: "title", Content: "cntnt", CreationDate: time.Now()},
			commentRepository:  repository.CustomCommentRepository(make([]model.Comment, 0)),
			postRepository:     repository.CustomPostRepository(make([]model.Post, 0)),
			expectedHttpStatus: 200,
			expectedHeader:     "application/json",
			expectedResponse:   RestJsonResponse{Message: "post id: 256 successfully added", Status: http.StatusOK},
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(tt *testing.T) {
			// GIVEN
			data, _ := json.Marshal(tc.post)
			req := httptest.NewRequest(http.MethodPost, "https://example.com/api/post/post", bytes.NewReader(data))
			w := httptest.NewRecorder()
			svc := RestApiService{&tc.postRepository, &tc.commentRepository}

			// WHEN
			handlePostAdd(&svc)(w, req)
			response := w.Result()
			body, _ := io.ReadAll(response.Body)
			var ackResponse RestJsonResponse
			err := json.Unmarshal(body, &ackResponse)
			if err != nil {
				t.Fail()
			}

			// THEN
			assert.Equal(t, tc.expectedHttpStatus, response.StatusCode)
			assert.Equal(t, tc.expectedHeader, response.Header.Get("Content-Type"))
			assert.Equal(t, tc.expectedResponse, ackResponse)
		})
	}
}
