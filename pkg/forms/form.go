package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

// to hold any validation errors for the form data
type Form struct {
	url.Values
	Errors errors
}

// to initialize a Form struct, it takes the form data as the parameter
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// to check if the required fields in the form are present
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" { // returns a slice of string with all white spaces removed
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// checks if it reaches max length or not
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (maximum is %d", d))
	}
}

// checks if it has specific permitted values
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid")
}

// returns true if there are no errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
