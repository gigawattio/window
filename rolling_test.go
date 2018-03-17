package window

import (
	"reflect"
	"testing"
)

func TestRolling(t *testing.T) {
	testCases := []struct {
		elements []string
		n        int
		expected [][]string
	}{
		{
			elements: nil,
			n:        0,
			expected: nil,
		},
		{
			elements: []string{},
			n:        0,
			expected: nil,
		},
		{
			elements: []string{"bail", "early"},
			n:        -1,
			expected: nil,
		},
		{
			elements: []string{"foo", "bar", "baz"},
			n:        0,
			expected: nil,
		},
		{
			elements: []string{"foo", "bar", "baz"},
			n:        1,
			expected: [][]string{
				[]string{"foo"},
				[]string{"bar"},
				[]string{"baz"},
			},
		},
		{
			elements: []string{"foo", "bar", "baz"},
			n:        1,
			expected: [][]string{
				[]string{"foo"},
				[]string{"bar"},
				[]string{"baz"},
			},
		},
		{
			elements: []string{"foo", "bar", "baz"},
			n:        2,
			expected: [][]string{
				[]string{"foo", "bar"},
				[]string{"bar", "baz"},
			},
		},
		{
			elements: []string{"foo", "bar", "baz"},
			n:        3,
			expected: [][]string{
				[]string{"foo", "bar", "baz"},
			},
		},
		{
			elements: []string{"foo", "bar", "baz", "boo"},
			n:        2,
			expected: [][]string{
				[]string{"foo", "bar"},
				[]string{"bar", "baz"},
				[]string{"baz", "boo"},
			},
		},
	}

	for i, testCase := range testCases {
		actual := Rolling(testCase.elements, testCase.n)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf("[i=%v] Actual rolling window result did not match expected value\n\tExpected: %+v\n\t  Actual: %+v", i, testCase.expected, actual)
		}
	}
}
