build:
	@go build -o bin/todolist cmd/main.go

run: build
	@./bin/todolist

test:
	@go test -v ./...