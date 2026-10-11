package testing

import (
	"testing"

	"example.com/hello/transport"
)

func TestQueueHandle(t *testing.T) {
	transport.HandleQueue(5)
}
