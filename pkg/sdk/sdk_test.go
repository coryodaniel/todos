package sdk_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/coryodaniel/todo/pkg/sdk"
	"github.com/coryodaniel/todo/pkg/todo"
)

func TestAddTodo(t *testing.T) {

}

func TestListTodos(t *testing.T) {
	wantItems := []todo.Item{
		{Title: "Wash dog"},
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assertEquals(t, req.URL.String(), "/api/todos/")

		// Send response to be tested
		resp, _ := json.Marshal(&wantItems)
		rw.Write(resp)
	}))

	defer mockServer.Close()

	api := sdk.API{mockServer.Client(), mockServer.URL}
	gotItems, _ := api.ListTodos()

	assertDeepEquals(t, &wantItems, gotItems)
}

func assertDeepEquals(t testing.TB, want any, got any) {
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected %v, got %v", want, got)
	}
}

func assertEquals(t testing.TB, want string, got string) {
	if want != got {
		t.Errorf("expected %v, got %v", want, got)
	}
}
