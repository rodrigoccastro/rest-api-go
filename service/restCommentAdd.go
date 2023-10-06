package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rodrigoccastro/rest-api-go/model"
)

func handleCommentAdd(svc *RestApiService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		var comment model.Comment

		if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}

		if err := svc.commentRepository.Insert(comment); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} 
		
		data, _ := json.Marshal(&RestJsonResponse{Message: fmt.Sprintf("comment id: %d successfully added", comment.Id), Status: http.StatusOK})
		if _, err := w.Write(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		
	}
}