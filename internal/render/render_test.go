package render

import (
	"net/http"
	"testing"

	"github.com/ppichugin/booking-for-breakfast/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)
	if result.Flash != "123" {
		t.Error("flash value 123 not found in session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplate()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = Template(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error writing template to browser")
	}

	err = Template(&ww, r, "none-existent.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("rendered template that doesn't exist")
	}
}

func TestNewTemplates(t *testing.T) {
	NewRenderer(app)
}

func TestCreateTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplate()
	if err != nil {
		t.Error(err)
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil
}
