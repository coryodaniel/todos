package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/coryodaniel/todo/internal/server"
	"github.com/coryodaniel/todo/pkg/todo"
)

type InMemoryTodoStore struct {
}

func (i *InMemoryTodoStore) GetTodo(id string) *todo.Item {
	return &todo.Item{}
}

func (i *InMemoryTodoStore) CreateTodo(params *todo.Item) (*todo.Item, error) {
	return nil, errors.New("not implemented")
}

func (i *InMemoryTodoStore) ListTodos() *[]todo.Item {
	return &[]todo.Item{}
}

func main() {
	addr := ":3333"
	fmt.Printf("Started todo server on %s. Args: %v", addr, os.Args[1:])

	server := &server.TodoServer{Store: &InMemoryTodoStore{}}
	log.Fatal(http.ListenAndServe(addr, server))
}
