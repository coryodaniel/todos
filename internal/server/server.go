package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/coryodaniel/todo/pkg/todo"
)

type TodoStore interface {
	GetTodo(id string) *todo.Item
	CreateTodo(params *todo.Item) (*todo.Item, error)
	ListTodos() *[]todo.Item
}

type TodoServer struct {
	Store TodoStore
}

func (t *TodoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		t.createTodo(w, r)
	case http.MethodGet:
		t.getTodo(w, r)
	}
}

func (t *TodoServer) getTodo(w http.ResponseWriter, r *http.Request) {
	todoId := strings.TrimPrefix(r.URL.Path, "/api/todos/")

	if todoId == "" {
		items := t.Store.ListTodos()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
		return
	}

	item := t.Store.GetTodo(todoId)

	if item != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(item)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (t *TodoServer) createTodo(w http.ResponseWriter, r *http.Request) {
	var params todo.Item
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)

	item, _ := t.Store.CreateTodo(&params)

	if item != nil {
		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(item)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}
