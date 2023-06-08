package sdk

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/coryodaniel/todo/pkg/todo"
)

type API struct {
	Client  *http.Client
	BaseURL string
}

func NewClient(url string) *API {
	return &API{
		Client:  http.DefaultClient,
		BaseURL: url,
	}
}

func (api *API) CreateTodo(title string) (*todo.Item, error) {
	params := map[string]string{"title": title}
	reqJson, _ := json.Marshal(params)
	reqBody := bytes.NewBuffer(reqJson)

	resp, err := api.Client.Post(api.BaseURL+"/api/todos/", "application/json", reqBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	item := todo.Item{}
	json.Unmarshal(body, &item)

	return &item, err

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
