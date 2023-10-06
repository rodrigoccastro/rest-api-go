package service

import (
	"encoding/json"
	"net/http"
)

func handlePostList(svc *RestApiService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		posts := svc.postRepository.GetAll(); 
		data, _ := json.Marshal(&posts)
		w.Write(data)
		return
	}
}