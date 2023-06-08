GOPATH?=${HOME}/go
ADDR?=http://localhost:3333/api/todos/

.PHONY: all
all:
	go fmt
	go test ./... -cover
	go build

docs:
	go install golang.org/x/tools/cmd/godoc@latest
	${GOPATH}/bin/godoc -http=localhost:6060

todos:
	curl -X POST ${ADDR} -H 'Content-Type: application/json' -d '{"title":"Wash car"}'
