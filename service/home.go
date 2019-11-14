package service

import (
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, req *http.Request) {
	u := struct {
		ID      string `json:"id"`
		Content string `json:"content"`
	}{ID: req.URL.Query().Get("id"), Content: req.URL.Query().Get("content")}

	templ, _ := template.ParseFiles("assets/templates/table.html")
	templ.Execute(w, u)
	w.Header().Set("Content-Type", "text/html")
}

func unknown(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 page not implemented", http.StatusNotImplemented)
}

func init() {
	http.HandleFunc("/api/form", homeHandler)
	http.HandleFunc("/unknown", unknown)
}
