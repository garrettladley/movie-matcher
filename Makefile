.PHONY:build
build:
	@go build -o bin/movie_matcher cmd/server/main.go
	@go build -o bin/omdb cmd/omdb/main.go
	@go build -o bin/pref_gen cmd/pref_gen/main.go

.PHONY:run
run: build
	@./bin/movie_matcher
	
.PHONY: install 
install:
	@go get ./...
	@go mod tidy
	@go mod download
