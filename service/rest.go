package service

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/rodrigoccastro/rest-api-go/repository"
)

type RestApiService struct {
	postRepository    *repository.PostRepository
	commentRepository *repository.CommentRepository
}

type RestJsonResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewRestApiService() RestApiService {
	return RestApiService{postRepository: repository.NewPostRepository(), commentRepository: repository.NewCommentRepository()}
}

func (svc *RestApiService) ServeContent(port int) error {
	portString := ":" + strconv.Itoa(port)
	svc.initializeHandlers()
	return http.ListenAndServe(portString, nil)
}

func (svc *RestApiService) initializeHandlers() {

	//posts
	http.HandleFunc("/api/post/add",  handlePostAdd(svc))
	http.HandleFunc("/api/post/find/", handlePostFind(svc))	
	http.HandleFunc("/api/post/list", handlePostList(svc))
	http.HandleFunc("/api/post/update",  handlePostUpdate(svc))

	//comments
	http.HandleFunc("/api/comment/add",  handleCommentAdd(svc))
	http.HandleFunc("/api/comment/list/", handleCommentListByPostId(svc))

}

func GetPathId(r *http.Request) string {
	arr_path := strings.Split(r.URL.Path, "/")
	return arr_path[len(arr_path)-1]
}

