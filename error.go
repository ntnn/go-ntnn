package ntnn

import (
	"fmt"
	"slices"
)

var globalIgnoreErrors = []error{nil}

// IgnoreError configures which errors should be ignored.
func IgnoreError(err error) {
	globalIgnoreErrors = append(globalIgnoreErrors, err)
}

func isIgnoredErr(err error, ignoreErrors ...error) bool {
	if slices.Contains(globalIgnoreErrors, err) {
		return true
	}
	return slices.Contains(ignoreErrors, err)
}

// Error returns false if the error is nil, otherwise it logs the error
// and returns true.
func Error(err error) bool {
	if isIgnoredErr(err) {
		return false
	}
	Logf("error: %v", err)
	return true
}

// ErrorFn executes the provided function and passes its error result to
// Error.
func ErrorFn(fn func() error) bool {
	return Error(fn())
}

// ErrorFnV is like ErrorFn but for functions that expect exactly one
// argument.
func ErrorFnV[T any](t T, fn func(T) error) bool {
	return Error(fn(t))
}

// Errorf returns false if the error is nil, otherwise it logs the
// provided format and arguments with the error appended and returns
// true.
func Errorf(err error, format string, args ...any) bool {
	if isIgnoredErr(err) {
		return false
	}
	Logf("%s: %v", fmt.Sprintf(format, args...), err)
	return true
}

// Panic does nothing if the error is nil, otherwise it logs the error
// and panics.
func Panic(err error) {
	if isIgnoredErr(err) {
		return
	}
	Logf("panic: %v", err)
	panic(err)
}

// PanicFn executes the provided function and passes its error result to
// Panic.
func PanicFn(fn func() error) {
	Panic(fn())
}

// Panicf does nothing if the error is nil, otherwise it logs the
// message and error.
func Panicf(message string, err error) {
	if isIgnoredErr(err) {
		return
	}
	Logf("%s: %v", message, err)
	panic(err)
}

// PanicfFn executes the provided function and passes its error result
// with the message to Panicf.
func PanicfFn(msg string, fn func() error) {
	Panicf(msg, fn())
}
