package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/coryodaniel/todo/pkg/todo"
)

type TodoHandler struct {
	Store TodoStore
}

func (t *TodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		t.create(w, r)
	case http.MethodGet:
		t.get(w, r)
	}
}

func (t *TodoHandler) get(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GET] %s\n", r.URL.Path)
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

func (t *TodoHandler) create(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)
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
