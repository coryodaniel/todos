package server

import (
	"embed"
	"log"
	"net/http"
)

//go:embed assets/js/* assets/node_modules/*
var assets embed.FS

func NewServer(addr string) error {
	log.Printf("Started todo server on %s.\n", addr)
	mux := http.NewServeMux()

	mux.Handle("/api/todos/", &TodoHandler{Store: NewMemoryStore()})

	mux.Handle("/assets/", assetsHandler())
	mux.Handle("/", &AppHandler{})

	return http.ListenAndServe(addr, mux)
}

func assetsHandler() http.Handler {
	fs := http.FileServer(http.FS(assets))
	return loggingHandler(http.StripPrefix("/", fs))
}

func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}
