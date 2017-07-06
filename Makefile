all:
	go build -o bin/go-ps-rest *.go
test:
	go test -v *.go
