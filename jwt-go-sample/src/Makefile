jwt-go-sample: *.go */*.go go.mod
	gofmt -w *.go */*.go
	go build -o jwt-go-sample main.go

.PHONY: init
init:
	go mod init jwt-go-sample

.PHONY: test
test:
	gofmt -w *.go
	go test
