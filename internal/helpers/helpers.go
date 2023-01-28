package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/ppichugin/booking-for-breakfast/internal/config"
)

var app *config.AppConfig

// NewHelpers sets up the application configuration for the helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error: ", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error, debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
