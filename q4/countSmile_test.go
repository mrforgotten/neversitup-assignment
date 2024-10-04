package smileface_test

import (
	"smileface"
	"testing"
)

type countSmileTestCase struct {
	input  []string
	expect int
}

func TestCountSmile(t *testing.T) {
	testCases := []countSmileTestCase{
		{
			input:  []string{":)", ";(", ";}", ":-D"},
			expect: 2,
		},
		{
			input:  []string{";D", ":-(", ":-)", ";~)"},
			expect: 3,
		},
		{
			input:  []string{";]", ":[", ";*", ":$", ";-D"},
			expect: 1,
		},
	}

	for _, tc := range testCases {
		got := smileface.CountSmile(tc.input)

		if tc.expect != got {
			t.Errorf("\nexpect: %v\ngot   : %v\ninput : %v", tc.expect, got, tc.input)
			continue
		}

	}
}
