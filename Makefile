GOPATH?=${HOME}/go
ADDR?=http://localhost:5555/api/todos/

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
	curl -X POST ${ADDR} -H 'Content-Type: application/json' -d '{"title":"Wash car"}'

start: build
	./todo server -addr=":5555"
