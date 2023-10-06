package main

import (
	"log"

	"github.com/rodrigoccastro/rest-api-go/bootstrap"
)

func main() {
	defaultPort := 8080
	if err := bootstrap.Init(defaultPort); err != nil {
		log.Fatalf("Service will be shutdown because error ocurred:  %+v", err.Error())
	}
}
