package randomstring

import (
	"fmt"
	"testing"
)

func Test_create_and_retrieve_user(t *testing.T) {
	expectedLength := 10
	var code string = String(expectedLength)
	fmt.Printf("Code generated: %s.\n", code)

	if expectedLength != len(code) {
		t.Errorf("Code length %d is not equals to required length %d.\n", len(code), expectedLength)
	}
}
