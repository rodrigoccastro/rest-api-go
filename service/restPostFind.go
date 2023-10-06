package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func handlePostFind(svc *RestApiService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		
		w.Header().Set("Content-Type", "application/json")
		str_id := GetPathId(r)
		id, err := strconv.ParseUint(str_id, 10, 64);

		if err != nil {
			data,_ := json.Marshal(&RestJsonResponse{Message: fmt.Sprintf("wrong id path variable: %s", str_id), Status: http.StatusBadRequest})
			w.Write(data)
			return
		}

		post, _ := svc.postRepository.GetById(id); 
		if post==nil {
			data, _ := json.Marshal(&RestJsonResponse{Message: fmt.Sprintf("Post with id: %d does not exist", id), Status: http.StatusBadRequest})
			w.Write(data)
			return
		}
		
		data, _ := json.Marshal(&post)
		w.Write(data)
		return
	}
}
