package pairs

import (
	"testing"
)

type testStruct struct {
	tests                  StructForPairs
	expected_value         bool
	expected_slice_of_ints []int64
}

var tests = []testStruct{
	{
		StructForPairs{[]int64{1, 2, 4, 9}, 8}, false, nil,
	},
	{
		StructForPairs{[]int64{1, 2, 4, 4}, 8}, true, []int64{4, 4},
	},
	{
		StructForPairs{[]int64{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 5}, 8}, true, []int64{3, 5},
	},
}

func TestGetPairs(t *testing.T) {
	for _, structure := range tests {
		success_value, slice_of_ints := GetPairsThatMatchesSum(structure.tests)
		if success_value != structure.expected_value {
			t.Error(
				"Should have returned", structure.expected_value,
				"instead of", success_value,
			)
		}
		if slice_of_ints == nil && structure.expected_slice_of_ints == nil {
			continue
		}
		if slice_of_ints == nil || structure.expected_slice_of_ints == nil {
			t.Fatal("For", structure.tests.values,
				"was expected", structure.expected_slice_of_ints,
				"instead of", slice_of_ints,
			)
		}
		for key := range slice_of_ints {
			if slice_of_ints[key] != structure.expected_slice_of_ints[key] {
				t.Error(
					"For", structure.tests.values,
					"was expected", structure.expected_slice_of_ints,
					"instead of", slice_of_ints,
				)
			}
		}
	}
}
