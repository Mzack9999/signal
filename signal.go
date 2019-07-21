package signal

import (
	"os"
	"os/signal"
	"syscall"
)

// Setup graceful exit with anonymous callback
func Setup(callback func()) {
	exitchan := make(chan os.Signal, 1)
	signal.Notify(exitchan, syscall.SIGTERM)
	signal.Notify(exitchan, syscall.SIGINT)

	go func() {
		<-exitchan
		callback()
	}()
}
