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

var (
	testDateUpdateRest = time.Date(2018, time.September, 16, 12, 0, 0, 0, time.UTC)
)

var validPostUpdateRest = model.Post{Id: 34, Title: "happy post", Content: "test content", CreationDate: testDateUpdateRest}

func TestPostUpdateRest(t *testing.T) {
	t.Run("UpdatePost", func(tt *testing.T) {
		// GIVEN
		post := model.Post{Id: 34, Title: "title", Content: "cntnt", CreationDate: time.Now()}
		data, _ := json.Marshal(post)
		req := httptest.NewRequest(http.MethodPost, "https://example.com/api/post/post", bytes.NewReader(data))
		w := httptest.NewRecorder()
		repoPost := repository.CustomPostRepository([]model.Post{validPostUpdateRest})
		repoComment := repository.CustomCommentRepository(make([]model.Comment, 0))
		svc := RestApiService{&repoPost, &repoComment}
		expectedResponse := RestJsonResponse{Message: "post id: 34 successfully updated", Status: http.StatusOK}

		// WHEN
		handlePostUpdate(&svc)(w, req)
		response := w.Result()
		body, _ := io.ReadAll(response.Body)
		var ackResponse RestJsonResponse
		err := json.Unmarshal(body, &ackResponse)
		if err != nil {
			t.Fail()
		}

		// THEN
		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, "application/json", response.Header.Get("Content-Type"))
		assert.Equal(t, expectedResponse, ackResponse)
	})
}
