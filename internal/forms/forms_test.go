package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	type fields struct {
		Values *http.Request
		Errors errors
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "POST", fields: fields{httptest.NewRequest("POST", "/api", nil), nil}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Form{
				Values: tt.fields.Values.Form,
				Errors: tt.fields.Errors,
			}
			if got := f.Valid(); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required filed missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_IsEmail(t *testing.T) {
	type fields struct {
		Email string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "email_valid", fields: fields{"test@test.com"}, want: ""},
		{name: "email_valid", fields: fields{"my@email.kz"}, want: ""},
		{name: "email_not_valid.", fields: fields{"test@testcom"}, want: "Invalid email address"},
		{name: "email_not_valid", fields: fields{"test@test"}, want: "Invalid email address"},
		{name: "email_not_valid", fields: fields{"test"}, want: "Invalid email address"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postedData := url.Values{}
			postedData.Add("email", tt.fields.Email)
			form := New(postedData)
			form.IsEmail("email")
			if got := form.Errors.Get("email"); got != tt.want {
				t.Errorf("IsEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForm_MinLength(t *testing.T) {
	type fields struct {
		Values url.Values
		Errors errors
	}
	type args struct {
		field  string
		value  string
		length int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"Length is greater than minimum", fields{}, args{"Name", "Surname", 2}, true},
		{"Length is less than minimum", fields{}, args{"Name", "Surname", 20}, false},
		{"Length is less than minimum", fields{}, args{"Name", "S", 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postedData := url.Values{}
			postedData.Add(tt.args.field, tt.args.value)
			f := New(postedData)
			if got := f.MinLength(tt.args.field, tt.args.length); got != tt.want {
				t.Errorf("MinLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForm_Has(t *testing.T) {
	type fields struct {
		Values url.Values
		Errors errors
	}
	type args struct {
		field string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"Field is not empty", fields{}, args{"Name", "Surname"}, true},
		{"Field is empty", fields{}, args{"Name", ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postedData := url.Values{}
			postedData.Add(tt.args.field, tt.args.value)
			f := New(postedData)
			if got := f.Has(tt.args.field); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}
