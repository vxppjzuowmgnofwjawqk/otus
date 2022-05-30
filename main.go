package main

import (
	"net/http"
	"otus/internal/router"
	"otus/internal/storage"
)

func main() {
	r := router.New(storage.New())
	http.ListenAndServe(":8080", r.RootHandler())
}
