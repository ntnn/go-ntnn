package ntnn

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"time"
)

// StartProfileServer starts the pprof server on an unused port and
// returns the address and a function to stop the server.
func StartProfileServer(profile string) (string, func()) {
	server := &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", UnusedPort()),
		Handler: pprof.Handler(profile),
	}
	go func() { Errorf(server.ListenAndServe(), "server exited") }()
	return server.Addr, func() { Errorf(server.Close(), "error closing http server") }
}

// StartProfileServerAndStall calls StartProfileServer, logs the address
// and stalls indefinitely.
func StartProfileServerAndStall(profile string) {
	addr, stop := StartProfileServer(profile)
	defer stop()
	Logf("server available at: %q", addr)
	time.Sleep(99999 * time.Minute)
}

// DumpProfileTraceSeconds is the default number of seconds to collect
// when collecting the trace profile.
var DumpProfileTraceSeconds = "10"

// DumpToFile starts the pprof server, fetches the given profile and
// dumps it to the specified file.
func DumpProfile(profile, pathPrefix string) {
	Log("starting server to dump")
	addr, stop := StartProfileServer(profile)
	defer stop()

	pathPrefix += "-" + profile
	baseAddr := fmt.Sprintf("http://%s/debug/pprof/", addr)

	switch profile {
	case "goroutine":
		DumpToFile(baseAddr+profile+"?debug=0", pathPrefix+"-debug0.out")
		DumpToFile(baseAddr+profile+"?debug=2", pathPrefix+"-debug2.out")
	case "trace":
		DumpToFile(baseAddr+profile+"?seconds="+DumpProfileTraceSeconds, pathPrefix+".out")
	default:
		DumpToFile(baseAddr+profile, pathPrefix+".out")
	}
}
