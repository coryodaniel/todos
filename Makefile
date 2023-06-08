GOPATH?=${HOME}/go
TODO_PORT?=5555
TODO_URL?=http://localhost:${TODO_PORT}/api/todos/

.PHONY: all
all: clean fmt test build

.PHONY: fmt
fmt:
	go fmt

.PHONY: test
test:
	go test ./... -cover

.PHONY: build
build:
	go build

.PHONY: clean
clean:
	rm -f ./todo

.PHONY: docs
docs:
	go install golang.org/x/tools/cmd/godoc@latest
	${GOPATH}/bin/godoc -http=localhost:6060

todos:
	./todo client new buy hat
	./todo client new mow yard
	./todo client new rake leaves
	./todo client new feed dog

start: build
	./todo server -addr=":${TODO_PORT}"
