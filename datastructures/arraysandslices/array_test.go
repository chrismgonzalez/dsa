package arraysandslices

import (
	"reflect"
	"testing"
)

func TestLength(t *testing.T) {
	//test the built-in length method
	numbers := [5]int{1, 2, 3, 4, 5}

	got := len(numbers)
	want := 5

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}


func TestSum(t *testing.T) {
	tests := []struct {
		input [4]int
		want int
	}{
		{input: [4]int{1, 2, 3, 4}, want: 10},
		{input: [4]int{0, 1, 2, 3}, want: 6},
		{input: [4]int{0, 0, 0, 0}, want: 0},
		{input: [4]int{}, want: 0},
	}


	for _, tc := range tests{
		got, _ := Sum(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %d, got: %d", tc.want, got)
		}
	}
}