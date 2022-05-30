package router

import (
	"net/http"
	"otus/internal/storage"
)

type router struct {
	rootHandler
}

func (r router) RootHandler() http.Handler {
	return r.rootHandler
}

func New(s storage.Storage) *router {
	return &router{
		rootHandler: newRootHandler(s),
	}
}
