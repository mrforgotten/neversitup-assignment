package main

import (
	"testing"
)

type oddIntTestCase struct {
	input  []int
	expect int
}

func TestOddInt(t *testing.T) {
	testCases := []oddIntTestCase{
		{
			input:  []int{1},
			expect: 1,
		},
		{
			input:  []int{1, 2, 1},
			expect: 2,
		},
		{
			input:  []int{0, 1, 0, 1, 0},
			expect: 0,
		},
		{
			input:  []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1},
			expect: 4,
		},
	}

	for _, tc := range testCases {
		got := OddInt(tc.input)

		if tc.expect != got {
			t.Errorf("\nexpect: %v\ngot   : %v\ninput : %v", tc.expect, got, tc.input)
			continue
		}

	}
}
