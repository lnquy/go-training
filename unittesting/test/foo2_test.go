package test

import (
	"fmt"
	"testing"
	"reflect"
	"gitlab.com/lnquy/go-training/unittesting/foo2"
)

// StubInputServicer is a foo struct using to test foo2.Reader methods
type StubInputServicer struct {
}

// Data implemented from foo2.InputServicer interface
func (s StubInputServicer) Data(url string) ([]byte, error) {
	if url != "" {
		return []byte(url), nil
	}
	return nil, fmt.Errorf("Error")
}

// TestFetchToGetByteArray tests when foo2.Reader.Fetch() return actual []byte
func TestFetchToGetBytes(test *testing.T) {
	reader := foo2.Reader{
		// Since StubInputServicer already implemented InputServicer interface
		// So it can be used here
		Servicer: StubInputServicer{},
	}

	actualResult, err := reader.Fetch("https://google.com")
	if err != nil {
		// Use Fatal/Fatalf to stop current test,
		// Error/Errorf will continue to run
		test.Fatalf("Expected no error, got %s", err)
	}

	// TODO
	expectedResult := []byte("define expected results for test cases (test table) here")
	if !reflect.DeepEqual(expectedResult, actualResult) {
		test.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

// TestFetchToGetByteArray tests when foo2.Reader.Fetch() return actual []byte
func TestFetchToReturnError(test *testing.T) {
	reader := foo2.Reader{
		Servicer: StubInputServicer{},
	}

	_, err := reader.Fetch("")
	if err == nil {
		test.Errorf("Expected error, got nil")
	}
}