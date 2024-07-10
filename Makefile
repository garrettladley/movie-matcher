.PHONY: install 
install:
	@go get ./...
	@go mod tidy
	@go mod download

.PHONY:build
build:
	@go build -o bin/movie_matcher main.go

.PHONY:run
run: build
	@./bin/movie_matcher