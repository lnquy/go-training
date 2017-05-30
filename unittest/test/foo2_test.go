// Black box testing

package test

import (
	"testing"
	"gitlab.com/lnquy/go-training/unittest/foo2"
	"errors"
	"reflect"
)

// StubInputServicer is a foo struct using to test foo2.Reader methods
type StubInputServicer struct {
}

// Data implemented from foo2.InputServicer interface
func (s StubInputServicer) GetData(url string) ([]byte, error) {
	if url != "" {
		return []byte(url), nil
	}
	return nil, errors.New("Error")
}

// TestFetchToGetByteArray tests when foo2.Reader.Fetch() return actual []byte
func TestFetchToGetBytes(test *testing.T) {
	reader := foo2.Reader{
		// Since StubInputServicer already implemented InputServicer interface
		// So it can be used here
		Servicer: StubInputServicer{},
	}

	// Test table
	testCases := []TestCase{
		{Input: "https://google.com", Output: "define expected results here"},
		{Input: "https://github.com", Output: "define expected results here"},
		{Input: "https://gitlab.com", Output: "define expected results here"},
		{Input: "https://medium.com", Output: "define expected results here"},
	}

	for _, tc := range testCases {
		actualResult, err := reader.Fetch(tc.Input.(string))
		if err != nil {
			// Use Fatal/Fatalf to stop current test,
			// Error/Errorf will continue to run
			test.Fatalf("foo2.Fetch(%s) = %s. Expected no error", tc.Input.(string), err)
		}

		if !reflect.DeepEqual([]byte(tc.Output.(string)), actualResult) {
			test.Errorf("foo2.Fetch(%s) = %v. Expected %v", tc.Input.(string),
				actualResult, []byte(tc.Output.(string)))
		}
	}
}

// TestFetchToGetByteArray tests when foo2.Reader.Fetch() return actual []byte
func TestFetchToReturnError(test *testing.T) {
	reader := foo2.Reader{
		Servicer: StubInputServicer{},
	}

	_, err := reader.Fetch("")
	if err == nil {
		test.Errorf("foo2.Fetch(%s) = nil. Expected error", "")
	}
}
