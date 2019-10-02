.PHONY: deps clean build

deps:
	go get -u ./...

clean:
	rm -rf ./declaration/declaration

build:
	GOOS=linux GOARCH=amd64 go build -o declaration/declaration ./declaration
