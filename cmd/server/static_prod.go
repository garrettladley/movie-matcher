//go:build !dev

package main

import (
	"embed"
	"net/http"
)

//go:embed public
var publicFS embed.FS // nolint:unused

// nolint:unused
func public() http.Handler {
	return http.FileServerFS(publicFS)
}

//go:embed deps
var depsFS embed.FS // nolint:unused

// nolint:unused
func deps() http.Handler {
	return http.FileServerFS(depsFS)
}

//go:embed images
var imagesFS embed.FS // nolint:unused

// nolint:unused
func images() http.Handler {
	return http.FileServerFS(imagesFS)
}
