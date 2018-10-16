package main

import (
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	type Tc struct {
		name          string
		input         string
		expectedMatch bool
		expectedWords []string
	}
	tcs := []Tc{
		{
			name:          "exists in dictionary no exact match",
			input:         "he",
			expectedMatch: false,
			expectedWords: []string{"hell", "hello", "helmet"},
		},
		{
			name:          "does not exist in dictionary",
			input:         "kg",
			expectedMatch: false,
			expectedWords: []string{},
		},
		{
			name:          "exists in dictionary and exact match",
			input:         "hell",
			expectedMatch: true,
			expectedWords: []string{"hello"},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(test *testing.T) {
			dict := LoadDict("./test_dict.txt")
			defer dict.Close()
			match, words := Find(tc.input, dict)
			if match != tc.expectedMatch {
				t.Errorf("%v match should exist for %v, got: %v, want: %v.", tc.name, tc.input, match, tc.expectedMatch)
			}
			if len(words) != len(tc.expectedWords) {
				t.Errorf("%v failed, got: %v, want: %v.", tc.name, len(words), len(tc.expectedWords))
			}
			if !reflect.DeepEqual(words, tc.expectedWords) {
				t.Errorf("%v failed, got: %v, want: %v.", tc.name, words, tc.expectedWords)
			}
		})
	}
}
