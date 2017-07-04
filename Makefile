all:
	go build -o bin/ps-rest *.go
test:
	go test -v *.go
