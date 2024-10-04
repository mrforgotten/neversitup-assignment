package main

import (
	"reflect"
	"testing"
)

type permutationTestCase struct {
	input  string
	expect []string
}

func TestPermutation(t *testing.T) {
	testCases := []permutationTestCase{
		{
			input:  "a",
			expect: []string{"a"},
		},
		{
			input:  "ab",
			expect: []string{"ab", "ba"},
		},
		{
			input:  "abc",
			expect: []string{"cba", "cab", "acb", "abc", "bac", "bca"},
		},
		{
			input: "abcd",
			expect: []string{
				"abcd",
				"abdc",
				"acbd",
				"acdb",
				"adbc",
				"adcb",
				"bacd",
				"badc",
				"bcad",
				"bcda",
				"bdac",
				"bdca",
				"cabd",
				"cadb",
				"cbad",
				"cbda",
				"cdab",
				"cdba",
				"dabc",
				"dacb",
				"dbac",
				"dbca",
				"dcab",
				"dcba",
			},
		},
		{
			input:  "aabb",
			expect: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"},
		},
	}
	for _, tc := range testCases {
		got := Permutation(tc.input)

		if len(tc.expect) != len(got) {
			t.Errorf("\nexpect: %v\ngot   : %v\nexpect: length %v\ngot   : length %v\ninput : %v", tc.expect, got, len(tc.expect), len(got), tc.input)
			continue
		}
		ansMap := make(map[string]int)
		expectMap := make(map[string]int)
		for i := range got {
			ansMap[got[i]]++
			expectMap[tc.expect[i]]++
		}
		if !reflect.DeepEqual(expectMap, ansMap) {

			t.Errorf("expect: %v\ngot   : %v\ninput :%v", tc.expect, got, tc.input)
		}
	}
}
