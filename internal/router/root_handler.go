package router

import (
	"net/http"
	"otus/internal/storage"
)

type rootHandler struct {
	tasksHandler
	taskHandler
}

func newRootHandler(s storage.Storage) rootHandler {
	return rootHandler{
		tasksHandler: newTasksHandler(s),
		taskHandler:  newTaskHandler(s),
	}
}

func (rh rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = cut(r.URL.Path)
	switch head {
	case "task":
		rh.taskHandler.ServeHTTP(w, r)
	case "tasks":
		rh.tasksHandler.ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}
