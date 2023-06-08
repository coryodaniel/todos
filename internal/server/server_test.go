package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coryodaniel/todo/internal/server"
	"github.com/coryodaniel/todo/pkg/todo"
)

func TestPOSTTodo(t *testing.T) {
	store := server.NewMemoryStore()

	todoServer := &server.TodoServer{Store: store}

	t.Run("it creates a todo", func(t *testing.T) {
		item := todo.Item{Title: "Mow yard"}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(&item)

		request, _ := http.NewRequest(http.MethodPost, "/api/todos/", b)
		response := httptest.NewRecorder()

		todoServer.ServeHTTP(response, request)

		if response.Code != http.StatusAccepted {
			t.Errorf("did not get correct status, got %d, want %d", response.Code, http.StatusAccepted)
		}

		result := todo.Item{}
		json.Unmarshal(response.Body.Bytes(), &result)

		want := item.Title
		got := result.Title

		if got != want {
			t.Errorf("wrong todo item returned, got %s want %s", got, want)
		}

		if len(*store.ListTodos()) != 1 {
			t.Errorf("got %d todos, wanted %d", len(*store.ListTodos()), 1)
		}
	})
}

func TestGETTodo(t *testing.T) {
	store := server.NewMemoryStore()
	store.CreateTodo(&todo.Item{Title: "Feed dog"})
	store.CreateTodo(&todo.Item{Title: "Wash dishes"})

	todoServer := &server.TodoServer{Store: store}

	t.Run("returns all todo items", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/todos/", nil)
		response := httptest.NewRecorder()

		todoServer.ServeHTTP(response, request)

		result := []map[string]interface{}{}
		json.Unmarshal(response.Body.Bytes(), &result)

		got := len(result)
		want := 2

		if got != want {
			t.Errorf("number of results is wrong, got %q want %q", got, want)
		}
	})

	t.Run("returns a todo item", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/todos/1", nil)
		response := httptest.NewRecorder()

		todoServer.ServeHTTP(response, request)

		result := map[string]interface{}{}
		json.Unmarshal(response.Body.Bytes(), &result)

		got := result["title"]
		want := "Feed dog"

		if got != want {
			t.Errorf("response body is wrong, got %q want %q", got, want)
		}
	})

	t.Run("returns an OK status when the record exists", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/todos/2", nil)
		response := httptest.NewRecorder()

		todoServer.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusOK

		if got != want {
			t.Errorf("did not get correct status, got %d, want %d", got, want)
		}
	})

	t.Run("returns 404 on missing todos", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/todos/1337", nil)
		response := httptest.NewRecorder()

		todoServer.ServeHTTP(response, request)

		if response.Code != http.StatusNotFound {
			t.Errorf("did not get correct status, got %d, want %d", response.Code, http.StatusNotFound)
		}
	})
}

func TestCreateAndListTodos(t *testing.T) {
	store := server.NewMemoryStore()
	todoServer := server.TodoServer{Store: store}

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

	got := len(result)
	want := 3

	if response.Code != http.StatusOK {
		t.Errorf("did not get correct status, got %d, want %d", response.Code, http.StatusOK)
	}

	if got != want {
		t.Errorf("number of results is wrong, got %d want %d", got, want)
	}
}
