package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/coryodaniel/todo/internal/server"
)

func main() {
	addr := ":3333"
	fmt.Printf("Started todo server on %s.\nConfig: %v\n", addr, os.Args[1:])

	store := server.NewMemoryStore()
	server := &server.TodoServer{Store: store}
	log.Fatal(http.ListenAndServe(addr, server))
}
