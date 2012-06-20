package errors

import (
	"errors"
	"fmt"
)

type RequestError struct {
	Method string
	Ierror error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("%s got %s", r.Method, r.Ierror)
}

var (
	ErrMethod = errors.New("Method isn't supported")
	kErrUnimplemented = errors.New("Unimplemented function")
)