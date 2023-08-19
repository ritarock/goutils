package array

import (
	"math"
	"reflect"
	"testing"
)

func assertDeepEqual[T Number](t *testing.T, got, want []T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func assert[T Number | bool](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want []int
	}{
		{
			name: "success sort",
			arg:  []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Sort(test.arg)
			assertDeepEqual(t, got, test.want)
		})
	}
}

func TestReverseSort(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want []int
	}{
		{
			name: "success reverse sort",
			arg:  []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ReverseSort(test.arg)
			assertDeepEqual(t, got, test.want)
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want int
	}{
		{
			name: "slice length = 0",
			arg:  []int{},
			want: 0,
		},
		{
			name: "slice length > 0",
			arg:  []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
			want: 5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Max(test.arg)
			assert(t, got, test.want)
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want int
	}{
		{
			name: "slice length = 0",
			arg:  []int{},
			want: 0,
		},
		{
			name: "slice length > 0",
			arg:  []int{5, 4, 3, 2, 1, 2, 3, 4, 5},
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Min(test.arg)
			assert(t, got, test.want)
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want int
	}{
		{
			name: "slice length = 0",
			arg:  []int{},
			want: 0,
		},
		{
			name: "slice length > 0",
			arg:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: 55,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Sum(test.arg)
			assert(t, got, test.want)
		})
	}
}

func TestMean(t *testing.T) {
	tests := []struct {
		name string
		arg  []float64
		want float64
	}{
		{
			name: "slice length = 0",
			arg:  []float64{},
			want: 0,
		},
		{
			name: "slice length > 0",
			arg:  []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: 5.5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Mean(test.arg)
			assert(t, got, test.want)
		})
	}
}

func TestMedian(t *testing.T) {
	tests := []struct {
		name string
		arg  []float64
		want float64
	}{
		{
			name: "slice length = 0",
			arg:  []float64{},
			want: 0,
		},
		{
			name: "slice length is even",
			arg:  []float64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34},
			want: 4,
		},
		{
			name: "slice length is odd",
			arg:  []float64{0, 1, 1, 2, 3, 5, 8, 13, 21},
			want: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Median(test.arg)
			assert(t, got, test.want)
		})
	}
}

func TestMode(t *testing.T) {
	tests := []struct {
		name string
		arg  []float64
		want []float64
	}{
		{
			name: "slice length = 0",
			arg:  []float64{},
			want: []float64{},
		},
		{
			name: "one mode",
			arg:  []float64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34},
			want: []float64{1},
		},
		{
			name: "two mode",
			arg:  []float64{0, 0, 1, 1, 2, 3, 5, 8, 13, 21, 34},
			want: []float64{0, 1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Mode(test.arg)
			got = Sort(got)
			assertDeepEqual(t, got, test.want)
		})
	}
}
func TestStandardDeviation(t *testing.T) {
	tests := []struct {
		name string
		arg  []float64
		want float64
	}{
		{
			name: "slice length = 0",
			arg:  []float64{},
			want: 0,
		},
		{
			name: "success standard deviation",
			arg:  []float64{55, 55, 60, 65, 65},
			want: 4.472135955,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := StandardDeviation(test.arg)
			got = math.Round((got * 10000000000)) / 10000000000
			assert(t, got, test.want)
		})
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
			name:  "slice length = 0",
			array: []int{},
			want:  false,
		},
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
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Some(test.array, test.f)
			assert(t, got, test.want)
		})
	}
}

func TestEvery(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		f     func(int) bool
		want  bool
	}{
		{
			name:  "slice length = 0",
			array: []int{},
			want:  false,
		},
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
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Every(test.array, test.f)
			assert(t, got, test.want)
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		f     func(int) bool
		want  []int
	}{
		{
			name:  "slice length = 0",
			array: []int{},
			want:  []int{},
		},
		{
			name:  "success filter",
			array: []int{1, 2, 3, 4, 5},
			f: func(v int) bool {
				return v >= 3
			},
			want: []int{3, 4, 5},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Filter(test.array, test.f)
			assertDeepEqual(t, got, test.want)
		})
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want []int
	}{
		{
			name: "success unique",
			arg:  []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Unique(test.arg)
			assertDeepEqual(t, got, test.want)
		})
	}
}
