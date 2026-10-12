package testing

import (
	"testing"

	"example.com/hello/concurrency"
)

func TestCurrency(t *testing.T) {
	concurrency.TwoPrint()
}
