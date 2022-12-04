package forms

import (
	"net/http"
	"net/url"
)

// Form creates a custom form struct, embeds url.Values object
type Form struct {
	url.Values
	errors
}

// New initializes a form structure
func New(data url.Values) *Form {
	return &Form{
		data,
		map[string][]string{},
	}
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}
