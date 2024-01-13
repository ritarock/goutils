package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	tests := []struct {
		array []int
		want  []int
	}{
		{
			array: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		got := Sort(test.array)
		assert.Equal(t, test.want, got)
	}
}

func TestReverseSort(t *testing.T) {
	tests := []struct {
		array []int
		want  []int
	}{
		{
			array: []int{1, 2, 3, 4, 5},
			want:  []int{5, 4, 3, 2, 1},
		},
	}

	for _, test := range tests {
		got := ReverseSort(test.array)
		assert.Equal(t, test.want, got)
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
				return v >= 3
			},
			want: true,
		},
		{
			name:  "return false",
			array: []int{1, 2, 3, 4, 5},
			f: func(v int) bool {
				return v == 0
			},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Some(test.array, test.f)
			assert.Equal(t, test.want, got)
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
				return v >= 1
			},
			want: true,
		},
		{
			name:  "return false",
			array: []int{1, 2, 3, 4, 5},
			f: func(v int) bool {
				return v >= 3
			},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Every(test.array, test.f)
			assert.Equal(t, test.want, got)
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
			name:  "succeed",
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
			assert.Equal(t, test.want, got)
		})
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  []int
	}{
		{
			name:  "slice length = 0",
			array: []int{},
			want:  []int{},
		},
		{
			name:  "succeed",
			array: []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		got := Unique(test.array)
		assert.Equal(t, test.want, got)
	}
}

func TestInclude(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		value int
		want  bool
	}{
		{
			name:  "slice length = 0",
			array: []int{},
			value: 0,
			want:  false,
		},
		{
			name:  "include value",
			array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			value: 5,
			want:  true,
		},
		{
			name:  "not include value",
			array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			value: 0,
			want:  false,
		},
	}

	for _, test := range tests {
		got := Include(test.array, test.value)
		assert.Equal(t, test.want, got)
	}
}
