package test

import (
	"testing"
	"fmt"
	"gitlab.com/lnquy/go-training/unittest/foo3"
)

type Dummy struct {
	foo3.InputServicer // Embedded interface
}

// Override from foo3.InputServicer interface
func (d Dummy) GetData(url string) ([]byte, error) {
	fmt.Println("GetData method from foo3_test.Reader")
	return []byte("Testing"), nil
}

// Run test for this function to verify what GetData() function has been called
func TestFetchFoo3(t *testing.T) {
	dummy := Dummy{}
	dummy.GetData("")

	// Other test stuffs...
}
