package home

import (
	"net/http"
	"templates"
)

// Handler handle GET requests to the Homepage
func Handler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "home", nil)
}
