package sdk

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/coryodaniel/todo/pkg/todo"
)

type API struct {
	Client  *http.Client
	BaseURL string
}

func (api *API) ListTodos() (*[]todo.Item, error) {
	resp, err := api.Client.Get(api.BaseURL + "/api/todos/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	items := []todo.Item{}
	json.Unmarshal(body, &items)

	return &items, err
}
