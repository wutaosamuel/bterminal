package utils

import (
	"fmt"
	"errors"
)

// errorString is a  trivial implementation of error
type errorString struct {
	s   string
	err error
}

// New for new error
func (e *errorString) New() *errorString {
	return &errorString{}
}

// Error for an error
func (e *errorString) Error() string {
	return e.s
}

// Errors for all error
func (e *errorString) Errors() string {
	return e.s + " \\\n" + e.err.Error()
}

// Err is custom error custommed from string
func Err(text string) error {
	return errors.New(text)
}

// Errs is to merge a err and a string
func Errs(text string, err error) error {
	errs := &errorString{text, err}
	return errors.New(errs.Errors())
}

////////////////// General ///////////////////

// CheckPanic check err != nil panic
func CheckPanic(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
