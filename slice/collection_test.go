package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		array []int
		want  []int
	}{
		{
			name:  "succeed: empty slise",
			array: []int{},
			want:  []int{},
		},
		{
			name:  "succeed",
			array: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := Sort(test.array)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestReverseSort(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		array []int
		want  []int
	}{
		{
			name:  "succeed: empty slise",
			array: []int{},
			want:  []int{},
		},
		{
			name:  "succeed",
			array: []int{1, 2, 3, 4, 5},
			want:  []int{5, 4, 3, 2, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := ReverseSort(test.array)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestSome(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		array []int
		f     func(int) bool
		want  bool
	}{
		{
			name:  "succeed: empty slise",
			array: []int{},
			f: func(i int) bool {
				return i >= 3
			},
			want: false,
		},
		{
			name:  "succeed: return true",
			array: []int{1, 2, 3, 4, 5},
			f: func(i int) bool {
				return i >= 3
			},
			want: true,
		},
		{
			name:  "succeed: return false",
			array: []int{1, 2, 3, 4, 5},
			f: func(i int) bool {
				return i == 0
			},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := Some(test.array, test.f)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestEvery(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		array []int
		f     func(int) bool
		want  bool
	}{
		{
			name:  "succeed: empty slise",
			array: []int{},
			f: func(i int) bool {
				return i >= 3
			},
			want: false,
		},
		{
			name:  "succeed: return true",
			array: []int{1, 2, 3, 4, 5},
			f: func(i int) bool {
				return i >= 1
			},
			want: true,
		},
		{
			name:  "succeed: return false",
			array: []int{1, 2, 3, 4, 5},
			f: func(i int) bool {
				return i > 1
			},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := Every(test.array, test.f)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestFilter(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		array []int
		f     func(int) bool
		want  []int
	}{
		{
			name:  "succeed: empty slise",
			array: []int{},
			f: func(i int) bool {
				return i >= 3
			},
			want: []int{},
		},
		{
			name:  "succeed",
			array: []int{1, 2, 3, 4, 5},
			f: func(i int) bool {
				return i >= 3
			},
			want: []int{3, 4, 5},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := Filter(test.array, test.f)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestUnique(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		array []int
		want  []int
	}{
		{
			name:  "succeed: empty slise",
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
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := Unique(test.array)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestInclude(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		array []int
		value int
		want  bool
	}{
		{
			name:  "succeed: empty slise",
			array: []int{},
			value: 0,
			want:  false,
		},
		{
			name:  "succeed",
			array: []int{1, 2, 3, 4, 5},
			value: 3,
			want:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := Include(test.array, test.value)
			assert.Equal(t, test.want, got)
		})
	}
}
