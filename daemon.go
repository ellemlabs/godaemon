package daemon

import (
	"os"
	"os/signal"
	"syscall"
)

var signalChannel chan os.Signal
var shutdown func()
func handleSignals() {
	for sig := range signalChannel {
		Notice(sig.String())
		switch sig {
		case syscall.SIGTERM:
		case syscall.SIGINT:
			Info("shutting down")
			shutdown()
		default:
			// do nothing
		}

	}
}

// don't use an interface or struct. this package is not object oritated.
// If a program uses this package, the program is assumed to be a daemon itself.
func Init(version string, shutdownFunc func()) {
	Info("starting " + version)

	// capture signals
	shutdown = shutdownFunc
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel)
	go handleSignals()

	Info("daemon initialized")
}