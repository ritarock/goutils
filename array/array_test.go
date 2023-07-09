package array

import (
	"reflect"
	"testing"
)

func assert[T Number | bool](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func assertDeepEqual[T Number | string](t *testing.T, got, want []T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
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
		got := Unique([]int{1, 2, 3, 3, 4, 4, 5, 5})
		want := []int{1, 2, 3, 4, 5}
		assertDeepEqual(t, got, want)
	})
	t.Run("string slice", func(t *testing.T) {
		got := Unique([]string{"a", "b", "c", "c", "d", "d", "d", "e", "e"})
		want := []string{"a", "b", "c", "d", "e"}
		assertDeepEqual(t, got, want)
	})
}

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "slice length > 0",
			args: []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
			want: 5,
		},
		{
			name: "slice length = 0",
			args: []int{},
			want: 0,
		},
	}
	for _, tc := range tests {
		got := Max(tc.args)
		want := tc.want
		assert(t, want, got)
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "slice length > 0",
			args: []int{5, 4, 3, 2, 1, 2, 3, 4, 5},
			want: 1,
		},
		{
			name: "slice length = 0",
			args: []int{},
			want: 0,
		},
	}
	for _, tc := range tests {
		got := Min(tc.args)
		want := tc.want
		assert(t, want, got)
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "slice length > 0",
			args: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: 55,
		},
		{
			name: "slice length = 0",
			args: []int{},
			want: 0,
		},
	}
	for _, tc := range tests {
		got := Sum(tc.args)
		want := tc.want
		assert(t, want, got)
	}
}

func TestMean(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want float64
	}{
		{
			name: "slice length > 0",
			args: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: 5.5,
		},
		{
			name: "slice length = 0",
			args: []int{},
			want: 0,
		},
	}
	for _, tc := range tests {
		got := Mean(tc.args)
		want := tc.want
		assert(t, want, got)
	}
}

func TestMedian(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want float64
	}{
		{
			name: "slice length = 0",
			args: []int{},
			want: 0,
		},
		{
			name: "slice length is odd",
			args: []int{1, 2, 3, 4, 5},
			want: 3,
		},
		{
			name: "slice length is even",
			args: []int{1, 2, 3, 4, 5, 6},
			want: float64(3.5),
		},
	}
	for _, tc := range tests {
		got := Median(tc.args)
		want := tc.want
		assert(t, want, got)
	}
}

func TestMode(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "slice length = 0",
			args: []int{},
			want: []int{0},
		},
		{
			name: "one mode",
			args: []int{1, 2, 2, 3, 3, 3, 4, 4, 5, 5},
			want: []int{3},
		},
		{
			name: "two mode",
			args: []int{1, 2, 2, 3, 3, 3, 4, 4, 5, 5, 5},
			want: []int{3, 5},
		},
	}
	for _, tc := range tests {
		got := Mode(tc.args)
		want := tc.want
		got = Sort(got)
		assertDeepEqual(t, want, got)
	}
}

func TestSome(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		f     func(int) bool
		want  bool
	}{
		{
			name:  "return true",
			array: []int{1, 2, 3, 4, 5},
			f: func(v int) bool {
				return v == 3
			},
			want: true,
		},
		{
			name:  "return false",
			array: []int{1, 2, 3, 4, 5},
			f: func(v int) bool {
				return v == 6
			},
			want: false,
		},
	}
	for _, tc := range tests {
		got := Some(tc.array, tc.f)
		want := tc.want
		assert(t, got, want)
	}
}

func TestEver(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		f     func(int) bool
		want  bool
	}{
		{
			name:  "return true",
			array: []int{1, 2, 3, 4, 5},
			f: func(v int) bool {
				return v > 0
			},
			want: true,
		},
		{
			name:  "return false",
			array: []int{1, 2, 3, 4, 5},
			f: func(v int) bool {
				return v > 3
			},
			want: false,
		},
	}
	for _, tc := range tests {
		got := Ever(tc.array, tc.f)
		want := tc.want
		assert(t, got, want)
	}
}

func TestFilter(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		got := Filter(
			[]int{1, 2, 3, 4, 5},
			func(v int) bool {
				return v >= 3
			},
		)
		want := []int{3, 4, 5}
		assertDeepEqual(t, got, want)
	})
}
