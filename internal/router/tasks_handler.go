package router

import (
	"encoding/json"
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
		data, err := json.Marshal(list)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
	case http.MethodPost:

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
