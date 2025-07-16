package ntnn

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
)

var dumpStackToFileLock sync.Mutex

func callers(skip int) []runtime.Frame {
	pc := make([]uintptr, 999)
	count := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc)
	ret := make([]runtime.Frame, count)

	i := 0
	for {
		frame, more := frames.Next()
		ret[i] = frame
		i += 1
		if !more {
			break
		}
	}

	return ret
}

// Callers returns all frames from the call this function as collected
// by runtime.Callers and runtime.CallersFrames as a slice.
func Callers() []runtime.Frame {
	return callers(3)
}

// LogCallers logs all frames from the call of this function as
// collected by runtime.Callers and runtime.CallersFrames using Log.
func LogCallers() {
	frames := callers(3)
	s := make([]string, len(frames)+1)

	s[0] = "callers:"
	for i, frame := range frames {
		s[i+1] = fmt.Sprintf("  :%s:%d:%s:", frame.File, frame.Line, frame.Function)
	}

	Log(strings.Join(s, "\n"))
}

// Stack returns the current stack trace as a string.
func Stack() string {
	buf := make([]byte, 16*1024)
	runtime.Stack(buf, false)
	return string(bytes.Trim(buf, "\x00")) + "\n\n"
}

// DumpStackToFile dumps the current stack trace to a file at the
// specified path.
//
// preStack can be left empty. If it is not empty it will be dumped to
// the file before the current stack trace. Useful when tracing origins
// and uses of a private function or callback.
//
// additionalInfo can be left empty. If it is not empty it will be
// written between the preStack and the current stack trace.
func DumpStackToFile(path, preStack, additionalInfo string) {
	curStack := Stack()

	dumpStackToFileLock.Lock()
	defer dumpStackToFileLock.Unlock()

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	Panic(err)
	defer Panic(f.Close())

	out := ""
	if preStack != "" {
		out += "preStack: " + preStack
	}

	if additionalInfo != "" {
		out += "additional: " + additionalInfo + "\n"
	}

	out += "curStack: " + curStack

	if _, err := f.WriteString(out); err != nil {
		Panic(err)
	}
}
