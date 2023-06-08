package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coryodaniel/todo/internal/server"
	"github.com/coryodaniel/todo/pkg/todo"
)

func TestCreateAndGetTodos(t *testing.T) {
	fmt.Printf("WTF")
	store := InMemoryTodoStore{}
	todoServer := server.TodoServer{Store: &store}

	item := todo.Item{Title: "Mow yard"}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(&item)

	postReq, _ := http.NewRequest(http.MethodPost, "/api/todos/", b)
	todoServer.ServeHTTP(httptest.NewRecorder(), postReq)

	postReq, _ = http.NewRequest(http.MethodPost, "/api/todos/", b)
	todoServer.ServeHTTP(httptest.NewRecorder(), postReq)

	postReq, _ = http.NewRequest(http.MethodPost, "/api/todos/", b)
	todoServer.ServeHTTP(httptest.NewRecorder(), postReq)

	getReq, _ := http.NewRequest(http.MethodGet, "/api/todos/", b)
	response := httptest.NewRecorder()
	todoServer.ServeHTTP(response, getReq)

	result := []map[string]interface{}{}
	json.Unmarshal(response.Body.Bytes(), &result)

	fmt.Printf("What is %v", result)

	got := len(result)
	want := 3

	if response.Code != http.StatusOK {
		t.Errorf("did not get correct status, got %d, want %d", response.Code, http.StatusOK)
	}

	if got != want {
		t.Errorf("number of results is wrong, got %q want %q", got, want)
	}
}
