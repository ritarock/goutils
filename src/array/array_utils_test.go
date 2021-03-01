package array

import (
	"reflect"
	"testing"
)

var array []int = []int{1, 2, 3}

func TestToFloat64(t *testing.T) {
	intArray := []int{1, 2, 3}
	result := ToFloat64(intArray)
	expect := []float64{1, 2, 3}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("reult = %v, want %v", result, expect)
	}
}

func TestMaxOfArray(t *testing.T) {
	result := MaxOfArray(ToFloat64(array))
	var expect float64 = 3

	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestMinOfArray(t *testing.T) {
	result := MinOfArray(ToFloat64(array))
	var expect float64 = 1

	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestSumOfArray(t *testing.T) {
	result := SumOfArray(ToFloat64(array))
	var expect float64 = 6

	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestAvarageOfArray(t *testing.T) {
	result := AvarageOfArray(ToFloat64(array))
	var expect float64 = 2

	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestSortArray(t *testing.T) {
	array := []int{3, 1, 2, 10, 22}
	result := SortArray(ToFloat64(array))
	expect := []float64{1, 2, 3, 10, 22}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("reult = %v, want %v", result, expect)
	}
}

func TestReverseSortArray(t *testing.T) {
	array := []int{3, 1, 2, 10, 22}
	result := ReverseSortArray(ToFloat64(array))
	expect := []float64{22, 10, 3, 2, 1}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("reult = %v, want %v", result, expect)
	}
}

func TestUniqueArray(t *testing.T) {
	array := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 5}
	result := UniqueArray(ToFloat64(array))
	expect := []float64{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("result = %v, want %v", result, expect)
	}
}
