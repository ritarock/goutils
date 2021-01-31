package array

import (
	"reflect"
	"testing"
)

func TestToFloat64(t *testing.T) {
	intArray := []int{1, 2, 3}
	result := ToFloat64(intArray)
	expect := []float64{1, 2, 3}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("reult = %v, want %v", result, expect)
	}
}

func TestMaxOfArray(t *testing.T) {
	array := []int{1, 2, 3}
	result := MaxOfArray(ToFloat64(array))
	var expect float64 = 3

	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestMinOfArray(t *testing.T) {
	array := []int{1, 2, 3}
	result := MinOfArray(ToFloat64(array))
	var expect float64 = 1

	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestSumOfArray(t *testing.T) {
	array := []int{1, 2, 3}
	result := SumOfArray(ToFloat64(array))
	var expect float64 = 6

	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestAvarageOfArray(t *testing.T) {
	array := []int{1, 2, 3}
	result := AvarageOfArray(ToFloat64(array))
	var expect float64 = 2

	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestSortArray(t *testing.T) {
	array := []int{3, 1, 2}
	result := SortArray(ToFloat64(array))
	expect := []float64{1, 2, 3}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("reult = %v, want %v", result, expect)
	}
}
