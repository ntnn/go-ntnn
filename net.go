package ntnn

import (
	"io"
	"net"
	"net/http"
	"os"
)

// UnusedPort returns a free TCP port on localhost.
func UnusedPort() int {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	Panic(err)
	port := l.Addr().(*net.TCPAddr).Port
	Panic(l.Close())
	return port
}

// DumpToFile fetches the content from the given address and writes it
// to the specified file.
func DumpToFile(addr, out string) {
	resp, err := http.Get(addr)
	Panic(err)
	defer PanicFn(resp.Body.Close)

	f, err := os.Create(out)
	Panic(err)
	defer PanicFn(f.Close)

	if _, err := io.Copy(f, resp.Body); err != nil {
		Errorf(err, "error writing request body to file")
	}
}
