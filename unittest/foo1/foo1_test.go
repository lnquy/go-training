// White box testing

package foo1

import (
	. "gitlab.com/lnquy/go-training/unittest/test"
	"testing" // Need to import testing package
)

// TestPlus tests the private plus(int, int) int function
func TestPlus(test *testing.T) {
	// Defines test cases (test table)
	testCases := []TestCase{
		{IntPair{0, 0}, 0},
		{IntPair{1, 0}, 1},
		{IntPair{5, 10}, 15},
		{IntPair{-12, 8}, -4},
		{IntPair{6, -20}, -14},
		{IntPair{-10, -10}, -20},
	}

	// Loop through test cases and compare returned value from plus() function with expected test case value
	for _, tc := range testCases {
		actualValue := plus(tc.Input.(IntPair).A, tc.Input.(IntPair).B)
		if tc.Output.(int) != actualValue {
			test.Errorf("foo1.plus(%d, %d) = %d. Expected %d", tc.Input.(IntPair).A, tc.Input.(IntPair).B,
				actualValue, tc.Output) // Provide more into here if need
		}
	}
}
