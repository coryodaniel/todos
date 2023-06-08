package sdk_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/coryodaniel/todo/pkg/sdk"
	"github.com/coryodaniel/todo/pkg/todo"
)

func TestAddTodo(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assertEquals(t, req.URL.String(), "/api/todos/")

		item := todo.Item{}
		body, _ := io.ReadAll(req.Body)
		json.Unmarshal(body, &item)
		item.ID = "1"

		// Send response to be tested
		resp, _ := json.Marshal(&item)
		rw.Write(resp)
	}))

	defer mockServer.Close()

	api := sdk.API{mockServer.Client(), mockServer.URL}
	gotItem, _ := api.CreateTodo("Buy hotdogs")

	assertEquals(t, gotItem.Title, "Buy hotdogs")
	assertEquals(t, gotItem.ID, "1")

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

	assertDeepEquals(t, gotItems, &wantItems)
}

func assertDeepEquals(t testing.TB, got any, want any) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func assertEquals(t testing.TB, got string, want string) {
	if got != want {
		t.Errorf("got %v, expected %v, ", got, want)
	}
}
