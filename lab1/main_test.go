package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewDualList(t *testing.T) {

	res := NewDualList()

	require.Nil(t, res.tail)
	require.Nil(t, res.head)
}

func TestDualList_Append(t *testing.T) {

	// Добавление в пустой список
	caseList1 := NewDualList()
	caseData1 := "data 1"
	caseList1.Append(caseData1)
	require.Equal(t, caseList1.head.data, caseData1)
	require.Equal(t, caseList1.tail.data, caseData1)

	// Корректная приязка трех элементов
	caseData2 := "data 2"
	caseData3 := "data 3"
	caseList1.Append(caseData2)
	caseList1.Append(caseData3)
	require.Equal(t, caseList1.tail.data, caseData1)
	require.Equal(t, caseList1.head.data, caseData3)
	require.Equal(t, caseList1.head.prev.data, caseData2)
	require.Equal(t, caseList1.tail.next.data, caseData2)
}

func TestDualList_ToStringForward(t *testing.T) {

	filledDualList := NewDualList()
	data1 := "data 1"
	data2 := "data 2"
	data3 := "data 3"
	filledDualList.Append(data1)
	filledDualList.Append(data2)
	filledDualList.Append(data3)

	cases := []struct {
		name     string
		input    *DualList
		expected string
	}{
		{
			name:     "Empty",
			input:    &DualList{},
			expected: "",
		},
		{
			name:     "Filled",
			input:    filledDualList,
			expected: "data 1 -> data 2 -> data 3",
		},
	}

	for _, tc := range cases {

		actual := tc.input.ToStringForward()
		require.Equal(t, tc.expected, actual)
	}
}

func TestDualList_ToStringBackward(t *testing.T) {

	filledDualList := NewDualList()
	data1 := "data 1"
	data2 := "data 2"
	data3 := "data 3"
	filledDualList.Append(data1)
	filledDualList.Append(data2)
	filledDualList.Append(data3)

	cases := []struct {
		name     string
		input    *DualList
		expected string
	}{
		{
			name:     "Empty",
			input:    &DualList{},
			expected: "",
		},
		{
			name:     "Filled",
			input:    filledDualList,
			expected: "data 3 -> data 2 -> data 1",
		},
	}

	for _, tc := range cases {

		actual := tc.input.ToStringBackward()
		require.Equal(t, tc.expected, actual)
	}
}

func TestDualList_DataByIndex(t *testing.T) {

	filledDualList := NewDualList()
	data1 := "data 1"
	data2 := "data 2"
	data3 := "data 3"
	filledDualList.Append(data1)
	filledDualList.Append(data2)
	filledDualList.Append(data3)

	cases := []struct {
		name       string
		inputList  *DualList
		inputIndex int
		expected   string
		err        error
	}{
		{
			name:       "Empty",
			inputList:  &DualList{},
			inputIndex: 2,
			expected:   "",
			err:        ErrSearchInEmptyList,
		},
		{
			name:       "First index",
			inputList:  filledDualList,
			inputIndex: 0,
			expected:   data1,
			err:        nil,
		},
		{
			name:       "Out of range",
			inputList:  filledDualList,
			inputIndex: 3,
			expected:   "",
			err:        ErrOutOfRangeIndex,
		},
		{
			name:       "Last index",
			inputList:  filledDualList,
			inputIndex: 2,
			expected:   data3,
			err:        nil,
		},
	}

	for _, tc := range cases {

		actual, actErr := tc.inputList.DataByIndex(tc.inputIndex)
		if actErr != nil {
			require.Equal(t, tc.err, actErr)
			continue
		}
		require.Equal(t, tc.expected, actual)
	}
}
