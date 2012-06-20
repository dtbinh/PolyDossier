package errors

import (
	"errors"
)

type RequestError struct {
	Method string
	ierror error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("%s got %s", r.Method, r.ierror)
}

var (
	ErrMethod = errors.New("Method isn't supported")
)