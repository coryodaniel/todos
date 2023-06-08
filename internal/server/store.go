package server

import (
	"strconv"

	"github.com/coryodaniel/todo/pkg/todo"
)

type TodoStore interface {
	GetTodo(id string) *todo.Item
	CreateTodo(params *todo.Item) (*todo.Item, error)
	ListTodos() *[]todo.Item
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		todos: map[string]todo.Item{},
	}
}

type MemoryStore struct {
	todos map[string]todo.Item
}

func (m *MemoryStore) GetTodo(id string) *todo.Item {
	item, ok := m.todos[id]
	if ok {
		return &item

	}

	return nil
}

func (m *MemoryStore) CreateTodo(params *todo.Item) (*todo.Item, error) {
	var item todo.Item = *params
	id := strconv.Itoa(len(m.todos) + 1)
	item.ID = id

	m.todos[id] = item
	return &item, nil
}

func (m *MemoryStore) ListTodos() *[]todo.Item {
	items := []todo.Item{}
	for _, item := range m.todos {
		items = append(items, item)
	}

	return &items
}
