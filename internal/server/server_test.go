package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/coryodaniel/todo/internal/server"
	"github.com/coryodaniel/todo/pkg/todo"
)

type StubTodoStore struct {
	todos map[string]todo.Item
}

func (s *StubTodoStore) GetTodo(id string) *todo.Item {
	item, ok := s.todos[id]
	if ok {
		return &item
	}

	return nil
}

func (s *StubTodoStore) CreateTodo(params *todo.Item) (*todo.Item, error) {
	var item todo.Item = *params
	id := strconv.Itoa(len(s.todos))
	item.ID = id

	s.todos[id] = item
	return &item, nil
}

func (s *StubTodoStore) ListTodos() *[]todo.Item {
	items := []todo.Item{}
	for _, item := range s.todos {
		items = append(items, item)
	}

	return &items
}

func TestPOSTTodo(t *testing.T) {
	store := StubTodoStore{
		map[string]todo.Item{},
	}

	todoServer := &server.TodoServer{Store: &store}

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

		if len(store.todos) != 1 {
			t.Errorf("got %d todos, wanted %d", len(store.todos), 1)
		}
	})
}

func TestGETTodo(t *testing.T) {
	store := StubTodoStore{
		map[string]todo.Item{
			"1": {Title: "Feed dog"},
			"2": {Title: "Wash dishes"},
		},
	}

	todoServer := &server.TodoServer{Store: &store}

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
