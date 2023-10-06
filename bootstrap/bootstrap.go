package bootstrap

import (
	"github.com/rodrigoccastro/rest-api-go/service"
)

func Init(port int) error {
	api := service.NewRestApiService()
	return api.ServeContent(port)
}
