package router

import (
	"net/http"
	"otus/internal/storage"
)

type tasksHandler struct {
	storage storage.Storage
}

func newTasksHandler(s storage.Storage) tasksHandler {
	return tasksHandler{
		storage: s,
	}
}

func (th tasksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		list := th.storage.GetTaskList(r.Context())
		writeJSON(w, &list)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
