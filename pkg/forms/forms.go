package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/http"
	"net/url"
	"strings"
)

// Form creates a custom form structure, embeds a url.Values objects
type Form struct {
	url.Values
	Errors errors
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if a field in the form has value
func (f *Form)Has(field string, r *http.Request) bool {
	x := r.FormValue(field)
	if x == "" {
		f.Errors.Add(field, "This field is required.")
		return false
	}

	return true
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field is required.")
		}
	}
}

func (f *Form) CheckLength(min int, max int, r *http.Request, fields ...string) bool {
	for _, field := range fields {
		x := r.FormValue(field)
		if len(x) > min && len(x) <= max {
			f.Errors.Add(field, fmt.Sprintf("%+v requries at least %+d and at most %+d characters.", field, min, max))
			return false
		}
	}

	return true
}

func (f *Form) IsEmail(field string) bool {
	return govalidator.IsEmail(field)
}