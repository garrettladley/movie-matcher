//go:build !dev
// +build !dev

package main

import (
	"embed"
	"net/http"
)

//go:embed public
var publicFS embed.FS

func public() http.Handler {
	return http.FileServerFS(publicFS)
}

//go:embed deps
var depsFS embed.FS

func deps() http.Handler {
	return http.FileServerFS(depsFS)
}
