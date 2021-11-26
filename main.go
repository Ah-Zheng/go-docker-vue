package main

import (
	"go-docker-vue/router"
	"log"
)

func main() {
	server := router.SetUpRouter()

	if err := server.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
