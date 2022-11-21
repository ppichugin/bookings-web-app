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

	fmt.Println("Starting application on port", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
