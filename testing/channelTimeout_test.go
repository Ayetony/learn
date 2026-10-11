package testing

import (
	"testing"

	"example.com/hello/transport"
)

func TestTimeout(t *testing.T) {
	ch := make(chan string, 1)
	transport.TimeoutHandle(ch)
}

func TestDeadLock(t *testing.T) {
	transport.NoCSPVailable()
}
