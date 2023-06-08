GOPATH?=${HOME}/go

.PHONY: all
all:
	go fmt
	go test ./...
	go build

docs:
	go install golang.org/x/tools/cmd/godoc@latest
	${GOPATH}/bin/godoc -http=localhost:6060
