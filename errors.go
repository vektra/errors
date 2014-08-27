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

// ErrorString is a trivial implementation of error.
type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return e.s
}

// HereError wraps another error with location information
type HereError struct {
	error
	pc  uintptr
	loc string
}

// Wrap an error with location information derived from the caller location
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

// Return a good string representation of the location and error
func (h *HereError) Error() string {
	return h.Location() + ": " + h.error.Error()
}

// Return the full path and line information for the location
func (h *HereError) FullLocation() string {
	return h.loc
}

// Return a short version of the location information
func (h *HereError) Location() string {
	lastSlash := strings.LastIndex(h.loc, "/")
	secondLastSlash := strings.LastIndex(h.loc[:lastSlash], "/")

	return h.loc[secondLastSlash+1:]
}

// Contains 2 errors, an updated error and a causing error
type CauseError struct {
	error
	cause error
}

// Wraps an error containing the information about what caused this error
func Cause(err error, cause error) *CauseError {
	return &CauseError{
		error: err,
		cause: cause,
	}
}

// Return the causing error
func (c *CauseError) Cause() error {
	return c.cause
}

// Contains an error and a stacktrace
type TraceError struct {
	error
	trace string
}

// Wraps an error with a stacktrace derived from the calling location
func Trace(err error) *TraceError {
	buf := make([]byte, 1024)
	sz := runtime.Stack(buf, false)

	return &TraceError{
		error: err,
		trace: string(buf[:sz]),
	}
}

// Return the stacktrace
func (t *TraceError) Trace() string {
	return t.trace
}
