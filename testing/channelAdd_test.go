package testing

import (
	"testing"

	"example.com/hello/transport"
)

func TestChannelAdd(t *testing.T) {
	transport.ExecChanAdd(10)
}
