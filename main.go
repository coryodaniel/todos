package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/coryodaniel/todo/internal/server"
	"github.com/coryodaniel/todo/pkg/todo"
)

type InMemoryTodoStore struct{}

func (i *InMemoryTodoStore) GetTodo(id string) *todo.Task {
	return &todo.Task{}
}

func (i *InMemoryTodoStore) CreateTodo(id string) {}

func main() {
	addr := ":3333"
	fmt.Printf("Started todo server on %s. Args: %v", addr, os.Args[1:])

	server := &server.TodoServer{Store: &InMemoryTodoStore{}}
	log.Fatal(http.ListenAndServe(addr, server))
}
