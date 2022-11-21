package main

import (
	"fmt"
	"github.com/ppichugin/booking-for-breakfast/pkg/config"
	"github.com/ppichugin/booking-for-breakfast/pkg/handlers"
	"github.com/ppichugin/booking-for-breakfast/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplate()
	if err != nil {
		log.Fatal("can not create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting application on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
