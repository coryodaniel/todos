package server

import (
	"embed"
	"log"
	"net/http"
)

//go:embed static/*
var assets embed.FS

func NewServer(addr string) error {
	log.Printf("Started todo server on %s.\n", addr)
	mux := http.NewServeMux()

	mux.Handle("/api/todos/", &TodoHandler{Store: NewMemoryStore()})

	// mux.Handle("/assets/", assetsHandler())
	mux.Handle("/", &AppHandler{})

	return http.ListenAndServe(addr, mux)
}

// func assetsHandler() http.Handler {
// 	getAllFilenames(&assets)

// 	fs := http.FileServer(http.FS(assets))

// 	return loggingHandler(http.StripPrefix("/", fs))
// }

// func loggingHandler(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println(r.Method, r.URL.Path)
// 		h.ServeHTTP(w, r)
// 	})
// }

// func getAllFilenames(efs *embed.FS) (files []string, err error) {
// 	if err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
// 		fmt.Printf("THe file name is: %s\n", path)
// 		if d.IsDir() {
// 			return nil
// 		}

// 		files = append(files, path)

// 		return nil
// 	}); err != nil {
// 		return nil, err
// 	}

// 	return files, nil
// }
