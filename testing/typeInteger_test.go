package testing

import (
	"fmt"
	"testing"

	"example.com/hello/object"
)

func TestTypeInteger(t *testing.T) {
	var a1 object.Integer = 1
	var a2 object.Integer = 9
	fmt.Printf("a1.Less(a2): %v\n", a1.Less(a2))
}
