Phony: build test deps run

PROJECT=socnet


test:
	go test -v ./$(PROJECT)

deps:
	go get -d golang.org/x/crypto/acme/autocert
	go get -u gonum.org/v1/gonum/...
	go get ./...

build: deps
	go build -o snet snet.go
	
run:
	go run snet.go
