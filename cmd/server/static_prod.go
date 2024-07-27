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

//go:embed htmx
var htmxFS embed.FS

func htmx() http.Handler {
	return http.FileServerFS(htmxFS)
}
