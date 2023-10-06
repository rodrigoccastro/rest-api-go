package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rodrigoccastro/rest-api-go/model"
)

func handlePostAdd(svc *RestApiService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		var post model.Post
		
		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}

		if err := svc.postRepository.Insert(post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} 
		
		data, _ := json.Marshal(&RestJsonResponse{Message: fmt.Sprintf("post id: %d successfully added", post.Id), Status: http.StatusOK})
		if _, err := w.Write(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}