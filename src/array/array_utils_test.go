package array

import (
	"reflect"
	"testing"
)

func TestToFloat64(t *testing.T) {
	intArr := []int{1, 2, 3}
	floatArr := []float64{1, 2, 3}
	intResult := ToFloat64(intArr)
	floatResult := ToFloat64(floatArr)
	expect := []float64{1, 2, 3}

	if !reflect.DeepEqual(intResult, expect) {
		t.Errorf("result = %v, want %v", intResult, expect)
	}
	if !reflect.DeepEqual(floatResult, expect) {
		t.Errorf("result = %v, want %v", floatResult, expect)
	}
}

func TestSumArray(t *testing.T) {
	intArr := []int{1, 2, 3}
	floatArr := []float64{1, 2, 3}
	intResult := SumArray(intArr)
	floatResult := SumArray(floatArr)
	var expect float64 = 6

	if intResult != expect {
		t.Errorf("intResult = %v, want = %v", intResult, expect)
	}
	if floatResult != expect {
		t.Errorf("floatResult = %v, want = %v", floatResult, expect)
	}
}

func TestMaxArray(t *testing.T) {
	intArr := []int{1, 2, 3}
	floatArr := []float64{1, 2, 3}
	intResult := MaxArray(intArr)
	floatResult := MaxArray(floatArr)
	var expect float64 = 3

	if intResult != expect {
		t.Errorf("intResult = %v, want = %v", intResult, expect)
	}
	if floatResult != expect {
		t.Errorf("floatResult = %v, want = %v", floatResult, expect)
	}
}

func TestMinArray(t *testing.T) {
	intArr := []int{1, 2, 3}
	floatArr := []float64{1, 2, 3}
	intResult := MinArray(intArr)
	floatResult := MinArray(floatArr)
	var expect float64 = 1

	if intResult != expect {
		t.Errorf("intResult = %v, want = %v", intResult, expect)
	}
	if floatResult != expect {
		t.Errorf("floatResult = %v, want = %v", floatResult, expect)
	}
}

func TestAvarageArray(t *testing.T) {
	intArr := []int{1, 2, 3}
	floatArr := []float64{1, 2, 3}
	intResult := AvarageArray(intArr)
	floatResult := AvarageArray(floatArr)
	var expect float64 = 2

	if intResult != expect {
		t.Errorf("intResult = %v, want = %v", intResult, expect)
	}
	if floatResult != expect {
		t.Errorf("floatResult = %v, want = %v", floatResult, expect)
	}
}
