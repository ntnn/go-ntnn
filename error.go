package ntnn

import "fmt"

func Error(err error) bool {
	if err == nil {
		return false
	}
	Logf("error: %v", err)
	return true
}

func Errorf(err error, format string, args ...any) bool {
	if err == nil {
		return false
	}
	Logf("%s: %v", fmt.Sprintf(format, args...), err)
	return true
}

func Panic(err error) {
	if err == nil {
		return
	}
	Logf("panic: %v", err)
	panic(err)
}

func PanicFn(fn func() error) {
	Panic(fn())
}

func Panicf(message string, err error) {
	if err == nil {
		return
	}
	Logf("%s: %v", message, err)
	panic(err)
}

func PanicfFn(msg string, fn func() error) {
	Panicf(msg, fn())
}
