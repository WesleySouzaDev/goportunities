.PHONY: default run build test docs clean
# variables
APP_NAME=goportunities

# Tasks
default: run

run:
	@swag init
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./...
docs:
	@swag init
clean:
	@rm -rf $(APP_NAME)
	@rm -rf ./docs

