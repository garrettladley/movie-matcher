//+build dev
//go:build dev
// +build dev

package main

import (
	"net/http"
	"os"
)

func public() http.Handler {
	return http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public")))
}

func htmx() http.Handler {
	return http.StripPrefix("/htmx/", http.FileServerFS(os.DirFS("htmx")))
}
