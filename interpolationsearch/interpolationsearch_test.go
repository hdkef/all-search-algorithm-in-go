package interpolationsearch

import (
	"fmt"
	"testing"
)

func TestInterpolationSearchAsc(t *testing.T) {
	testTable := []struct {
		Input struct {
			Arr    []int
			Target int
		}
		Expected int
	}{
		{
			Input: struct {
				Arr    []int
				Target int
			}{
				Arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				Target: 10,
			},
			Expected: 9,
		},
		{
			Input: struct {
				Arr    []int
				Target int
			}{
				Arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				Target: 16,
			},
			Expected: -1,
		},
		{
			Input: struct {
				Arr    []int
				Target int
			}{
				Arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				Target: 4,
			},
			Expected: 3,
		},
		{
			Input: struct {
				Arr    []int
				Target int
			}{
				Arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				Target: -1,
			},
			Expected: -1,
		},
	}

	for i, v := range testTable {
		t.Run(fmt.Sprintf("test asc no %d", i), func(t *testing.T) {
			actual := InterpolationSearchAsc(v.Input.Arr, v.Input.Target)
			if actual != v.Expected {
				t.Errorf("expected %d got %d", v.Expected, actual)
			}
		})
	}
}
