package array

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		got := Sort([]int{5, 4, 3, 2, 1})
		want := []int{1, 2, 3, 4, 5}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("float64 slice", func(t *testing.T) {
		got := Sort([]float64{5, 4, 3, 2, 1})
		want := []float64{1, 2, 3, 4, 5}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestReverseSort(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		got := ReverseSort([]int{1, 2, 3, 4, 5})
		want := []int{5, 4, 3, 2, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("float64 slice", func(t *testing.T) {
		got := ReverseSort([]float64{1, 2, 3, 4, 5})
		want := []float64{5, 4, 3, 2, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestUnique(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		got := Unique([]int{1, 2, 3, 4, 5, 5, 5})
		want := []int{1, 2, 3, 4, 5}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("string slice", func(t *testing.T) {
		got := Unique([]string{"a", "b", "c", "d", "e", "e", "e"})
		want := []string{"a", "b", "c", "d", "e"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestMaxOfArray(t *testing.T) {
	got := MaxOfArray([]int{1, 2, 3, 4, 5})
	want := 5

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMinOfArray(t *testing.T) {
	got := MinOfArray([]int{1, 2, 3, 4, 5})
	want := 1

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumOfArray(t *testing.T) {
	got := SumOfArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	want := 55

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestAvarageOfArray(t *testing.T) {
	got := AvarageOfArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	want := 5.5

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
