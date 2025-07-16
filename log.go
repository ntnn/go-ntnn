package ntnn

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	Marker     = "###>"
	EnableLogs = true

	fileLock  sync.Mutex
	LogToFile = ""
)

func init() {
	if LogToFile != "" {
		// Ensure the log file exists
		f, err := os.OpenFile(LogToFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		Panicf("error opening LogToFile", err)
		Panic(f.Close())
	}
}

func printer(msg string) {
	if !EnableLogs {
		return
	}

	if !strings.HasSuffix(msg, "\n") {
		msg += "\n"
	}

	if LogToFile == "" {
		_, err := fmt.Print(Marker + " " + msg)
		Panic(err)
		return
	}

	fileLock.Lock()
	defer fileLock.Unlock()

	f, err := os.OpenFile(LogToFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	Panic(err)
	defer Panic(f.Close())

	_, err = f.WriteString(msg)
	Panic(err)
}

func Log(msg string) {
	printer(msg)
}

func Logf(format string, args ...any) {
	printer(fmt.Sprintf(format, args...))
}
