package main

import (
	"testing"
)

var testCases map[string]int

func init() {
	testCases = map[string]int{
		"@0x1a8510f2:0x1a8510f2.space": 3,
		"@tr_slimey:matrix.org": 8,
		"abcdefg": 5,
		"qwerty": 1,
		"qwertz": 8,
		"matrix": 8,
	}
}

func TestGetColorNumberFromHash(t *testing.T) {
	for testcase, expected := range testCases {
		hash := GetMxidHash(testcase)
		if result := GetColorNumberFromHash(hash); result+1 != expected {
			t.Errorf("got %d, want %d from %s", result, expected, testcase)
		}
	}
}
