package test

// Global struct holding all test cases
type TestCase struct {
	Input  interface{}
	Output interface{}
}

// IntPair holds input arguments for test cases of plus() function
type IntPair struct {
	A int
	B int
}

// IntTuple holds input arguments for test cases of foo1.PublicPlus() function
type IntTuple struct {
	A int
	B int
	C int
}
