package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeJSON(w http.ResponseWriter, value interface{}) {
	data, err := json.Marshal(value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(data)
	_, err = w.Write(data)
	if err != nil {
		fmt.Println(err)
	}
}
