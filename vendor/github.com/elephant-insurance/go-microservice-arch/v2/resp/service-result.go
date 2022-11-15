package resp

import (
	"bytes"
	"strings"
)

// ServiceResult wraps a result from the model to be rendered in the controller
type ServiceResult struct {
	ID     string `json:"id,omitempty" jsonapi:"-"`
	errors map[string]*error
}

// Renderable objects can contain a list of errors to be returned with the response
type Renderable interface {
	Errors() []JsonApiError
}

// Errors ...
func (r ServiceResult) Errors() []JsonApiError {
	errors := []JsonApiError{}
	var err JsonApiError

	for k, v := range r.errors {
		key := dashBeforeCaps(k)
		if v != nil {
			thisError := *v
			err = JsonApiError{Detail: thisError.Error(), Status: "422"}
			err.Source = "data/attributes/" + key
			errors = append(errors, err)
		}
	}

	return errors
}

// dashBeforeCaps inserts a dash before each capital letter
func dashBeforeCaps(str string) string {
	runeToInsert := rune('-')
	buf := bytes.NewBufferString("")
	for i, v := range str {
		if i > 0 && v >= 'A' && v <= 'Z' {
			buf.WriteRune(runeToInsert)
		}
		buf.WriteRune(v)
	}

	return strings.ToLower(buf.String())
}
