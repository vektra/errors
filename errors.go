package errors

import (
	"fmt"
	"runtime"
	"strings"
)

// Imported functionality from the stdlib errors package

// New returns an error that formats as the given text.
func New(text string) *ErrorString {
	return &ErrorString{text}
}

// Created a new error formated according to the fmt rules.
func Format(f string, val ...interface{}) error {
	return fmt.Errorf(f, val...)
}

// errorString is a trivial implementation of error.
type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return e.s
}

type HereError struct {
	error
	pc  uintptr
	loc string
}

func Here(orig error) *HereError {
	pc, file, line, ok := runtime.Caller(1)

	if ok {
		return &HereError{
			error: orig,
			pc:    pc,
			loc:   fmt.Sprintf("%s:%d", file, line),
		}
	}

	return &HereError{error: orig}
}

func (h *HereError) Error() string {
	return h.Location() + ": " + h.error.Error()
}

func (h *HereError) FullLocation() string {
	return h.loc
}

func (h *HereError) Location() string {
	lastSlash := strings.LastIndex(h.loc, "/")
	secondLastSlash := strings.LastIndex(h.loc[:lastSlash], "/")

	return h.loc[secondLastSlash+1:]
}

type CauseError struct {
	error
	cause error
}

func Cause(err error, cause error) *CauseError {
	return &CauseError{
		error: err,
		cause: cause,
	}
}

func (c *CauseError) Cause() error {
	return c.cause
}

type TraceError struct {
	error
	trace string
}

func Trace(err error) *TraceError {
	buf := make([]byte, 1024)
	sz := runtime.Stack(buf, false)

	return &TraceError{
		error: err,
		trace: string(buf[:sz]),
	}
}

func (t *TraceError) Trace() string {
	return t.trace
}
