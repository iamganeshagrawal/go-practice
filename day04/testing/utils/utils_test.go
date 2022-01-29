package utils

import "testing"

type testCase struct {
	name     string
	input    int32
	expected bool
	actual   bool
}

func Test_IsPrime(t *testing.T) {
	cases := []testCase{
		{name: "Testing_Prime: 31", input: 31, expected: true},
		{name: "Testing_Prime: 61", input: 61, expected: true},
		{name: "Testing_Prime: 71", input: 71, expected: true},
		{name: "Testing_Prime: 73", input: 73, expected: true},
		{name: "Testing_Prime: 64", input: 64, expected: false},
		{name: "Testing_Prime: 83", input: 83, expected: true},
		{name: "Testing_Prime: 12", input: 12, expected: false},
		{name: "Testing_Prime: 0", input: 0, expected: false},
		{name: "Testing_Prime: 2", input: 2, expected: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.actual = IsPrime(tc.input)
			if tc.actual != tc.expected {
				t.Errorf("Checking %d is Prime, expected %v, got %v", tc.input, tc.expected, tc.actual)
			}
		})
	}
}
