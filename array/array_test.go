package array

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	assert := func(t *testing.T, got interface{}, want interface{}) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("int", func(t *testing.T) {
		got := Sort([]int{5, 4, 3, 2, 1})
		want := []int{1, 2, 3, 4, 5}
		assert(t, got, want)
	})
	t.Run("float64", func(t *testing.T) {
		got := Sort([]float64{5, 4, 3, 2, 1})
		want := []float64{1, 2, 3, 4, 5}
		assert(t, got, want)
	})
	t.Run("others", func(t *testing.T) {
		got := Sort([]string{"1", "2", "3", "4", "5"})
		want := []string{"1", "2", "3", "4", "5"}
		assert(t, got, want)
	})
}

func TestReverseSort(t *testing.T) {
	assert := func(t *testing.T, got interface{}, want interface{}) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("int", func(t *testing.T) {
		got := ReverseSort([]int{1, 2, 3, 4, 5})
		want := []int{5, 4, 3, 2, 1}
		assert(t, got, want)
	})
	t.Run("float64", func(t *testing.T) {
		got := ReverseSort([]float64{1, 2, 3, 4, 5})
		want := []float64{5, 4, 3, 2, 1}
		assert(t, got, want)
	})
	t.Run("others", func(t *testing.T) {
		got := ReverseSort([]string{"1", "2", "3", "4", "5"})
		want := []string{"1", "2", "3", "4", "5"}
		assert(t, got, want)
	})
}

func TestUniqueArray(t *testing.T) {
	assert := func(t *testing.T, got interface{}, want interface{}) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("int", func(t *testing.T) {
		got := Unique([]int{1, 1, 2, 3, 4, 5, 5})
		want := []int{1, 2, 3, 4, 5}
		assert(t, got, want)
	})
	t.Run("float64", func(t *testing.T) {
		got := Unique([]float64{1, 1, 2, 3, 4, 5, 5})
		want := []float64{1, 2, 3, 4, 5}
		assert(t, got, want)
	})
	t.Run("string", func(t *testing.T) {
		got := Unique([]string{"a", "a", "b", "c", "d", "e", "e"})
		want := []string{"a", "b", "c", "d", "e"}
		assert(t, got, want)
	})
}

func TestMaxOfArray(t *testing.T) {
	assert := func(t *testing.T, got interface{}, want interface{}) {
		t.Helper()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("int", func(t *testing.T) {
		got := MaxOfArray([]int{1, 2, 3, 4, 5})
		want := 5
		assert(t, got, want)
	})
	t.Run("float64", func(t *testing.T) {
		got := MaxOfArray([]float64{1, 2, 3, 4, 5})
		want := 5.0
		assert(t, got, want)
	})
	t.Run("others", func(t *testing.T) {
		got := MaxOfArray([]string{"1", "2", "3", "4", "5"})
		want := 0
		assert(t, got, want)
	})
}

func TestMinOfArray(t *testing.T) {
	assert := func(t *testing.T, got interface{}, want interface{}) {
		t.Helper()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("int", func(t *testing.T) {
		got := MinOfArray([]int{1, 2, 3, 4, 5})
		want := 1
		assert(t, got, want)
	})
	t.Run("float64", func(t *testing.T) {
		got := MinOfArray([]float64{1, 2, 3, 4, 5})
		want := 1.0
		assert(t, got, want)
	})
	t.Run("others", func(t *testing.T) {
		got := MinOfArray([]string{"1", "2", "3", "4", "5"})
		want := 0
		assert(t, got, want)
	})
}

func TestSumOfArray(t *testing.T) {
	assert := func(t *testing.T, got interface{}, want interface{}) {
		t.Helper()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("int", func(t *testing.T) {
		got := SumOfArray([]int{1, 2, 3, 4, 5})
		want := 15
		assert(t, got, want)
	})
	t.Run("float64", func(t *testing.T) {
		got := SumOfArray([]float64{1, 2, 3, 4, 5})
		want := 15.0
		assert(t, got, want)
	})
	t.Run("others", func(t *testing.T) {
		got := SumOfArray([]string{"1", "2", "3", "4", "5"})
		want := 0
		assert(t, got, want)
	})
}

func TestAvarageOfArray(t *testing.T) {
	assert := func(t *testing.T, got interface{}, want interface{}) {
		t.Helper()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("int", func(t *testing.T) {
		got := AvarageOfArray([]int{1, 2, 3, 4, 5})
		want := 3.0
		assert(t, got, want)
	})
	t.Run("float64", func(t *testing.T) {
		got := AvarageOfArray([]float64{1, 2, 3, 4, 5})
		want := 3.0
		assert(t, got, want)
	})
	t.Run("int", func(t *testing.T) {
		got := AvarageOfArray([]string{"1", "2", "3", "4", "5"})
		want := 0.0
		assert(t, got, want)
	})
}
