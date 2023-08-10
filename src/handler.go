package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type Handler struct {
}

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	files := []string{
		"src/ui/html/base.go.tmpl",
		"src/ui/html/partials/navbar.go.tmpl",
		"src/ui/html/partials/footer.go.tmpl",
		"src/ui/html/home.go.tmpl",
	}

	template, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("Template Parsing Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = template.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println("Template Execution Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CreateSnippetHanlder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("create Snippet"))
}

func (h *Handler) ViewSnippetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}
