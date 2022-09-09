package errors

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"net/http"
)

func NewHTTPError(code int, field, detail string) *HTTPError {
	return &HTTPError{
		Errors: map[string][]string{
			field: {detail},
		},
		Code: code,
	}
}

// HTTPError is an HTTP conduitError.
type HTTPError struct {
	Errors map[string][]string `json:"errors"`

	Code int `json:"-"`
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTPError code: %d", e.Code)
}

func FromError(err error) *HTTPError {
	if err == nil {
		return nil
	}
	if se := new(errors.Error); errors.As(err, &se) {
		return NewHTTPError(http.StatusUnprocessableEntity, se.Reason, se.Message)
	}
	return NewHTTPError(http.StatusUnprocessableEntity, "internal", "error")
}
