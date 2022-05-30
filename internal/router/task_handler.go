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
		var task models.Task
		jsonDecoder := json.Decoder(r.Body)
		err := jsonDecoder.Decode(&task)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		th.storage.CreateTask(r.Context(), task)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
