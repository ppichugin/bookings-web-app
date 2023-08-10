package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"

	"github.com/ppichugin/booking-for-breakfast/internal/config"
	"github.com/ppichugin/booking-for-breakfast/internal/driver"
	"github.com/ppichugin/booking-for-breakfast/internal/handlers"
	"github.com/ppichugin/booking-for-breakfast/internal/helpers"
	"github.com/ppichugin/booking-for-breakfast/internal/models"
	"github.com/ppichugin/booking-for-breakfast/internal/render"
)

const portNumber = ":8080"

var (
	app      config.AppConfig
	session  *scs.SessionManager
	infoLog  *log.Logger
	errorLog *log.Logger
)

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()

	fmt.Println("Starting application on port", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	// what am I going to put into a session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	// change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // ssl not yet
	app.Session = session

	// connect to Database
	log.Println("Connecting to Database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=pg password=pg")
	if err != nil {
		log.Fatal("Can not connect to Database!")
	}
	log.Println("Connected to database!")

	tc, err := render.CreateTemplate()
	if err != nil {
		log.Fatal("can not create template cache")
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
