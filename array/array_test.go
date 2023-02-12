package array

import (
	"reflect"
	"testing"
)

func assertDeepEqual[T Number | string](t *testing.T, got, want []T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertNumber[T Number](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSort(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		got := Sort([]int{5, 4, 3, 2, 1})
		want := []int{1, 2, 3, 4, 5}
		assertDeepEqual(t, got, want)
	})
	t.Run("float slice", func(t *testing.T) {
		got := Sort([]float64{5, 4, 3, 2, 1})
		want := []float64{1, 2, 3, 4, 5}
		assertDeepEqual(t, got, want)
	})
}

func TestReverseSort(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		got := ReverseSort([]int{1, 2, 3, 4, 5})
		want := []int{5, 4, 3, 2, 1}
		assertDeepEqual(t, got, want)
	})
	t.Run("float slice", func(t *testing.T) {
		got := ReverseSort([]float64{1, 2, 3, 4, 5})
		want := []float64{5, 4, 3, 2, 1}
		assertDeepEqual(t, got, want)
	})
}

func TestUnique(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		got := Unique([]int{1, 2, 3, 4, 5, 5, 5})
		want := []int{1, 2, 3, 4, 5}
		assertDeepEqual(t, got, want)
	})
	t.Run("string slice", func(t *testing.T) {
		got := Unique([]string{"a", "b", "c", "d", "e", "e", "e"})
		want := []string{"a", "b", "c", "d", "e"}
		assertDeepEqual(t, got, want)
	})
}

func TestMaxOfArray(t *testing.T) {
	got := MaxOfArray([]int{1, 2, 3, 4, 5})
	want := 5
	assertNumber(t, got, want)
}

func TestMinOfArray(t *testing.T) {
	got := MinOfArray([]int{5, 4, 3, 2, 1})
	want := 1
	assertNumber(t, got, want)
}

func TestSumOfArray(t *testing.T) {
	got := SumOfArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	want := 45
	assertNumber(t, got, want)
}

func TestMeanOfArray(t *testing.T) {
	got := MeanOfArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	want := float64(5.5)
	assertNumber(t, got, want)
}

func TestMedian(t *testing.T) {
	t.Run("odd", func(t *testing.T) {
		got := MeanOfArray([]int{1, 2, 3, 4, 5})
		want := float64(3)
		assertNumber(t, got, want)
	})
	t.Run("even", func(t *testing.T) {
		got := MeanOfArray([]int{1, 2, 3, 4, 5, 6})
		want := float64(3.5)
		assertNumber(t, got, want)
	})
}
