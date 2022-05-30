package router

import (
	"encoding/json"
	"net/http"
	"otus/internal/models"
	"otus/internal/storage"
)

type taskHandler struct {
	storage storage.Storage
}

func newTaskHandler(s storage.Storage) taskHandler {
	return taskHandler{
		storage: s,
	}
}

func (th taskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		th.handlerCreateTask(w, r)
	case http.MethodDelete:
		th.handleDeleteTask(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (th taskHandler) handlerCreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	respTask := th.storage.CreateTask(r.Context(), task)
	writeJSON(w, &respTask)
}

func (th taskHandler) handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	id, _ := cut(r.URL.Path)
	th.storage.DeleteTask(r.Context(), id)
}
