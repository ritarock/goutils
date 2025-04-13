package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		array []int
		want  int
	}{
		{
			name:  "succeed: empty slise",
			array: []int{},
			want:  0,
		},
		{
			name:  "succeed",
			array: []int{1, 2, 3, 4, 5},
			want:  5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := Max(test.array)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMin(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		array []int
		want  int
	}{
		{
			name:  "succeed: empty slise",
			array: []int{},
			want:  0,
		},
		{
			name:  "succeed",
			array: []int{1, 2, 3, 4, 5},
			want:  1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := Min(test.array)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestSum(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		array []int
		want  int
	}{
		{
			name:  "succeed: empty slise",
			array: []int{},
			want:  0,
		},
		{
			name:  "succeed",
			array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:  55,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := Sum(test.array)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestAverage(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		array []int
		want  float64
	}{
		{
			name:  "succeed: empty slise",
			array: []int{},
			want:  0,
		},
		{
			name:  "succeed",
			array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:  5.5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := Average(test.array)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMedian(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		array []int
		want  float64
	}{
		{
			name:  "succeed: empty slise",
			array: []int{},
			want:  0,
		},
		{
			name:  "succeed: slice length is odd",
			array: []int{1, 1, 2, 3, 5, 8, 13, 21, 34},
			want:  5,
		},
		{
			name:  "succeed: slice length is even",
			array: []int{1, 1, 2, 3, 5, 8, 13, 21},
			want:  4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := Median(test.array)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMode(t *testing.T) {
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
			name:  "succeed: return one value",
			array: []int{1, 1, 2, 3, 4, 5, 6, 7},
			want:  []int{1},
		},
		{
			name:  "succeed: return two value",
			array: []int{1, 1, 2, 3, 4, 5, 6, 7, 7},
			want:  []int{1, 7},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := Mode(test.array)
			assert.Equal(t, test.want, got)
		})
	}
}
