package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	cases := []struct {
		name          string
		input         []int
		conditionFunc func(a int, b int) bool
		expected      []int
	}{
		{
			name:     "Empty",
			input:    []int{},
			expected: []int{},
		},
		{
			name:  "Default asc",
			input: []int{8, 7, 9, 1, 2, 6, 4, 5, 3, 0},
			conditionFunc: func(a int, b int) bool {
				return a < b
			},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "Default desc",
			input: []int{8, 7, 9, 1, 2, 6, 4, 5, 3, 0},
			conditionFunc: func(a int, b int) bool {
				return a > b
			},
			expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, tc := range cases {

		act := BubbleSort(tc.input, tc.conditionFunc)
		require.Equal(t, tc.expected, act)
	}
}

func TestShakerSort(t *testing.T) {

	cases := []struct {
		name          string
		input         []int
		conditionFunc func(a int, b int) bool
		expected      []int
	}{
		{
			name:     "Empty",
			input:    []int{},
			expected: []int{},
		},
		{
			name:  "Default asc",
			input: []int{8, 7, 9, 1, 2, 6, 4, 5, 3, 0},
			conditionFunc: func(a int, b int) bool {
				return a < b
			},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "Default desc",
			input: []int{8, 7, 9, 1, 2, 6, 4, 5, 3, 0},
			conditionFunc: func(a int, b int) bool {
				return a > b
			},
			expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, tc := range cases {

		act := ShakerSort(tc.input, tc.conditionFunc)
		require.Equal(t, tc.expected, act)
	}
}

func TestOddEvenSort(t *testing.T) {

	cases := []struct {
		name          string
		input         []int
		conditionFunc func(a int, b int) bool
		expected      []int
	}{
		{
			name:     "Empty",
			input:    []int{},
			expected: []int{},
		},
		{
			name:  "Default asc",
			input: []int{8, 7, 9, 1, 2, 6, 4, 5, 3, 0},
			conditionFunc: func(a int, b int) bool {
				return a < b
			},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "Default desc",
			input: []int{8, 7, 9, 1, 2, 6, 4, 5, 3, 0},
			conditionFunc: func(a int, b int) bool {
				return a > b
			},
			expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, tc := range cases {

		act := OddEvenSort(tc.input, tc.conditionFunc)
		require.Equal(t, tc.expected, act)
	}
}

func TestCombSort(t *testing.T) {

	cases := []struct {
		name          string
		input         []int
		conditionFunc func(a int, b int) bool
		expected      []int
	}{
		{
			name:     "Empty",
			input:    []int{},
			expected: []int{},
		},
		{
			name:  "Default asc",
			input: []int{8, 7, 9, 1, 2, 6, 4, 5, 3, 0},
			conditionFunc: func(a int, b int) bool {
				return a < b
			},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "Default desc",
			input: []int{8, 7, 9, 1, 2, 6, 4, 5, 3, 0},
			conditionFunc: func(a int, b int) bool {
				return a > b
			},
			expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, tc := range cases {

		act := CombSort(tc.input, tc.conditionFunc)
		require.Equal(t, tc.expected, act)
	}
}

func TestInsertSort(t *testing.T) {

	cases := []struct {
		name          string
		input         []int
		conditionFunc func(a int, b int) bool
		expected      []int
	}{
		{
			name:     "Empty",
			input:    []int{},
			expected: []int{},
		},
		{
			name:  "Default asc",
			input: []int{8, 7, 9, 1, 2, 6, 4, 5, 3, 0},
			conditionFunc: func(a int, b int) bool {
				return a < b
			},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "Default desc",
			input: []int{8, 7, 9, 1, 2, 6, 4, 5, 3, 0},
			conditionFunc: func(a int, b int) bool {
				return a > b
			},
			expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, tc := range cases {

		act := InsertSort(tc.input, tc.conditionFunc)
		require.Equal(t, tc.expected, act)
	}
}

func TestBinarySearch(t *testing.T) {

	cases := []struct {
		name        string
		input       []int
		searchItem  int
		expectedPos int
		err         error
	}{
		{
			name:        "Empty",
			input:       []int{},
			searchItem:  10,
			expectedPos: -1,
			err:         ErrIncorrectBinarySearchParams,
		},
		{
			name:        "Sorted",
			input:       []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			searchItem:  1,
			expectedPos: 1,
			err:         nil,
		},
		{
			name:        "Retry",
			input:       []int{1, 1, 1, 1, 1},
			searchItem:  1,
			expectedPos: 0,
			err:         nil,
		},
	}
	for _, tc := range cases {

		conditionFunc := func(a int, b int) bool {
			return a < b
		}
		pos, err := BinarySearch(tc.input, tc.searchItem, 0, len(tc.input)-1, conditionFunc)
		require.Equal(t, tc.err, err)
		require.Equal(t, tc.expectedPos, pos)
	}
}

func TestInsertBinarySort(t *testing.T) {

	cases := []struct {
		name          string
		input         []int
		conditionFunc func(a int, b int) bool
		expected      []int
	}{
		{
			name:     "Empty",
			input:    []int{},
			expected: []int{},
		},
		{
			name:  "Default asc",
			input: []int{8, 7, 9, 1, 2, 6, 4, 5, 3, 0},
			conditionFunc: func(a int, b int) bool {
				return a < b
			},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "Default desc",
			input: []int{8, 7, 9, 1, 2, 6, 4, 5, 3, 0},
			conditionFunc: func(a int, b int) bool {
				return a > b
			},
			expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, tc := range cases {

		act := InsertBinarySort(tc.input, tc.conditionFunc)
		require.Equal(t, tc.expected, act)
	}
}

func TestPairInsertSort(t *testing.T) {

	cases := []struct {
		name          string
		input         []int
		conditionFunc func(a int, b int) bool
		expected      []int
	}{
		{
			name:     "Empty",
			input:    []int{},
			expected: []int{},
		},
		{
			name:  "Default asc",
			input: []int{8, 7, 9, 1, 2, 6, 4, 5, 3, 0},
			conditionFunc: func(a int, b int) bool {
				return a < b
			},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		//{
		//	name:  "Default desc",
		//	input: []int{8, 7, 9, 1, 2, 6, 4, 5, 3, 0},
		//	conditionFunc: func(a int, b int) bool {
		//		return a > b
		//	},
		//	expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		//},
	}
	for _, tc := range cases {

		act := PairInsertSort(tc.input, tc.conditionFunc)
		require.Equal(t, tc.expected, act)
	}
}
