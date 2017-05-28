package test

import (
	"gitlab.com/lnquy/go-training/unittesting/foo1"
	"testing"
)

func TestPublicPlus(test *testing.T) {
	testCases := []TestCase{
		{IntTuple{0, 0, 0}, 0},
		{IntTuple{0, 5, 10}, 15},
		{IntTuple{9, 1, 4}, 14},
		{IntTuple{-1, 0, 0}, -1},
		{IntTuple{-10, -20, 10}, -20},
		{IntTuple{-100, -50, -100}, -250},
	}

	for _, tc := range testCases {
		actualValue := foo1.PublicPlus(tc.Input.(IntTuple).A, tc.Input.(IntTuple).B, tc.Input.(IntTuple).C)
		if tc.Output != actualValue {
			test.Errorf("Expected %d, got %d", tc.Output, actualValue)
		}
	}
}
