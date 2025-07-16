package ntnn

import "fmt"

// Error returns false if the error is nil, otherwise it logs the error
// and returns true.
func Error(err error) bool {
	if err == nil {
		return false
	}
	Logf("error: %v", err)
	return true
}

// Errorf returns false if the error is nil, otherwise it logs the
// provided format and arguments with the error appended and returns
// true.
func Errorf(err error, format string, args ...any) bool {
	if err == nil {
		return false
	}
	Logf("%s: %v", fmt.Sprintf(format, args...), err)
	return true
}

// Panic does nothing if the error is nil, otherwise it logs the error
// and panics.
func Panic(err error) {
	if err == nil {
		return
	}
	Logf("panic: %v", err)
	panic(err)
}

// PanicFn executed the provided function and passes its error result to
// Panic.
func PanicFn(fn func() error) {
	Panic(fn())
}

// Panicf does nothing if the error is nil, otherwise it logs the
// message and error.
func Panicf(message string, err error) {
	if err == nil {
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
