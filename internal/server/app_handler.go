package server

import (
	_ "embed"
	"net/http"
)

//go:embed ui/index.html
var indexPage []byte

type AppHandler struct{}

func (a *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write(indexPage)
}
