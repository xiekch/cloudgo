package service

import (
	"net/http"

	"encoding/json"
)

func apiTestHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(struct {
		ID      string `json:"id"`
		Content string `json:"content"`
	}{ID: "8675309", Content: "Hello from Go!"})
}

func init() {
	http.HandleFunc("/api/test", apiTestHandler)
}
