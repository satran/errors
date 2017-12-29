package errors

import (
	"fmt"
	"runtime"
)

type Error struct {
	message string
	caused  error // stores the original error
	file    string
	line    int
}

func WithCaused(err error, message string) *Error {
	return new(message, 3).CausedBy(err)
}

func New(message string) *Error {
	return new(message, 3)
}

func new(message string, count int) *Error {
	file, line := caller(count)
	return &Error{
		message: message,
		file:    file,
		line:    line,
	}
}

func (e *Error) CausedBy(err error) *Error {
	e.caused = err
	return e
}

func caller(count int) (string, int) {
	var ok bool
	_, file, line, ok := runtime.Caller(count)
	if !ok {
		file = "???"
		line = 0
	}
	return file, line
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Debug() string {
	msg := e.message
	if e.caused != nil {
		msg = e.caused.Error()
	}
	return fmt.Sprintf("%s:%d %s\n%s", e.file, e.line, msg, e.Error())
}
