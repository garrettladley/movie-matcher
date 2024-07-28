.PHONY:build
build: gen-css gen-templ
	@go build -tags dev -o bin/movie_matcher cmd/server/main.go
	@go build -o bin/omdb cmd/omdb/main.go
	@go build -o bin/pref_gen cmd/pref_gen/main.go

.PHONY:build-prod
build-prod: gen-css gen-templ
	@go build -tags prod -o bin/movie_matcher cmd/server/main.go
	@go build -o bin/omdb cmd/omdb/main.go
	@go build -o bin/pref_gen cmd/pref_gen/main.go

.PHONY:run
run: build
	@./bin/movie_matcher
	
.PHONY: install 
install: install-templ gen-templ
	@go get ./...
	@go mod tidy
	@go mod download
	@mkdir -p cmd/server/htmx
	@wget -q -O cmd/server/htmx/htmx.min.js.gz https://unpkg.com/htmx.org@1.9.12/dist/htmx.min.js.gz
	@gunzip -f cmd/server/htmx/htmx.min.js.gz
	@npm install -D daisyui
	@npm install -D flowbite
	@mkdir -p cmd/server/public/
	@cp -r node_modules/flowbite/dist/flowbite.min.js cmd/server/public/
	@cp -r node_modules/apexcharts/dist/apexcharts.min.js cmd/server/public/
	@npm install -D apexcharts
	@npm install -D tailwindcss


NODE_BIN := ./node_modules/.bin

.PHONY: gen-css
gen-css:
	@$(NODE_BIN)/tailwindcss build -i internal/views/css/app.css -o cmd/server/public/styles.css --minify

.PHONY: watch-css
watch-css:
	@$(NODE_BIN)/tailwindcss -i internal/views/css/app.css -o cmd/server/public/styles.css --minify --watch 

.PHONY: install-templ
install-templ:
	@go install github.com/a-h/templ/cmd/templ@latest

.PHONY: gen-templ
gen-templ:
	@templ generate

.PHONY: watch-templ
watch-templ:
	@templ generate --watch --proxy=http://127.0.0.1:8000

