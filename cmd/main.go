package main

import (
	"log"

	"github.com/aidarov11/todo-app"
	"github.com/aidarov11/todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error uccured while running http server: %s", err.Error())
	}
}
