package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSliceLength(t *testing.T) {
	//test the built-in length method
	numbers := []int{1, 2, 3, 4, 5}

	got := len(numbers)
	want := 5

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}


func TestSumSlice(t *testing.T) {
	tests := []struct {
		input []int
		want int
	}{
		{input: []int{}, want: 0},
		{input: []int{0, 0, 0, 0}, want: 0},
		{input: []int{1, 2, 3, 4}, want: 10},
		{input: []int{0, 1, 2, 3, 4, 5}, want: 15},
		{input: []int{0, 1, 2, 3, 4, 5, 6}, want: 21},
		{input: []int{0, 1, 3, 5, 7, 9, 11}, want: 36},
	
	}


	for _, tc := range tests{
		got, _ := SumSlice(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %d, got: %d", tc.want, got)
		}
	}
}