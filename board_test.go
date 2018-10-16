package main

import (
	"reflect"
	"testing"
)

func TestGetSurrounding(t *testing.T) {
	type Tc struct {
		name     string
		input    Pos
		rowSize  int
		expected []Pos
	}
	tcs := []Tc{
		{
			name:    "first position",
			input:   Pos{0, 0},
			rowSize: 4,
			expected: []Pos{
				Pos{0, 1},
				Pos{1, 0},
				Pos{1, 1},
			},
		},
		{
			name:    "mid top position",
			input:   Pos{0, 1},
			rowSize: 4,
			expected: []Pos{
				Pos{0, 0},
				Pos{0, 2},
				Pos{1, 0},
				Pos{1, 1},
				Pos{1, 2},
			},
		},
		{
			name:    "mid position",
			input:   Pos{1, 1},
			rowSize: 4,
			expected: []Pos{
				Pos{0, 0},
				Pos{0, 1},
				Pos{0, 2},
				Pos{1, 0},
				Pos{1, 2},
				Pos{2, 0},
				Pos{2, 1},
				Pos{2, 2},
			},
		},
		{
			name:    "end of grid position",
			input:   Pos{3, 3},
			rowSize: 4,
			expected: []Pos{
				Pos{2, 2},
				Pos{2, 3},
				Pos{3, 2},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			results := []Pos{}
			results = GetSurroundingPos(tc.input, tc.rowSize)
			if reflect.DeepEqual(results, tc.expected) == false {
				t.Errorf("%v failed, got: %d, want: %d.", tc.name, results, tc.expected)
			}
		})
	}
}

func TestCharsToGrid(t *testing.T) {
	type Tc struct {
		name     string
		input    []string
		gridSize int
		expected [][]string
	}
	tcs := []Tc{
		{
			name:     "4 by grid",
			input:    []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"},
			gridSize: 4,
			expected: [][]string{
				[]string{"a", "b", "c", "d"},
				[]string{"e", "f", "g", "h"},
				[]string{"i", "j", "k", "l"},
				[]string{"m", "n", "o", "p"},
			},
		},
		{
			name:     "3 by grid",
			input:    []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
			gridSize: 3,
			expected: [][]string{
				[]string{"a", "b", "c"},
				[]string{"d", "e", "f"},
				[]string{"g", "h", "i"},
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			results := [][]string{}
			results = CharsToGrid(tc.input, tc.gridSize)
			if len(results) != tc.gridSize {
				t.Errorf("%v length failed, got: %d, want: %d.", tc.name, len(results), tc.gridSize)
			}
			if reflect.DeepEqual(results, tc.expected) == false {
				t.Errorf("%v failed, got: %d, want: %d.", tc.name, results, tc.expected)
			}
		})
	}

}
