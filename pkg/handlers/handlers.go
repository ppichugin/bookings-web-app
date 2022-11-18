package handlers

import (
	"github.com/ppichugin/booking-for-breakfast/pkg/render"
	"net/http"
)

// Home is the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}