.PHONY: default run build test docs clean

#Variables
APP_NAME=go-opportunities

#Task
default: run-with-docs

run:
	@go run main.go
run-with-docs:
	@swag init --parseDependency --parseInternal --parseDepth 1
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init --parseDependency --parseInternal --parseDepth 1
clean:
	@rm -rf $(APP_NAME)
	@rm -rf ./docs