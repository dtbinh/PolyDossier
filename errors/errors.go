package errors

import (
	"errors"
	"fmt"
)

type RequestError struct {
	Action  string
	Method  string
	Problem string
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("[%s]<%s> -> %s", r.Action, r.Method, r.Problem)
}

var (
	ErrMethod        = errors.New("Method isn't supported")
	ErrUnimplemented = errors.New("Unimplemented function")
)
