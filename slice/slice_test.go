package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  int
	}{
		{
			name:  "array length == 0",
			array: []int{},
			want:  0,
		},
		{
			name:  "succeed",
			array: []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
			want:  5,
		},
	}

	for _, test := range tests {
		got := Max(test.array)
		assert.Equal(t, test.want, got)
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  int
	}{
		{
			name:  "array length == 0",
			array: []int{},
			want:  0,
		},
		{
			name:  "succeed",
			array: []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
			want:  1,
		},
	}

	for _, test := range tests {
		got := Min(test.array)
		assert.Equal(t, test.want, got)
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  int
	}{
		{
			name:  "array length == 0",
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
		got := Sum(test.array)
		assert.Equal(t, test.want, got)
	}
}

func TestMean(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  float64
	}{
		{
			name:  "array length == 0",
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
		got := Mean(test.array)
		assert.Equal(t, test.want, got)
	}
}

func TestMedian(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  float64
	}{
		{
			name:  "array length == 0",
			array: []int{},
			want:  0,
		},
		{
			name:  "slice length is odd",
			array: []int{1, 1, 2, 3, 5, 8, 13, 21, 34},
			want:  5,
		},
		{
			name:  "slice length is even",
			array: []int{1, 1, 2, 3, 5, 8, 13, 21},
			want:  4,
		},
	}

	for _, test := range tests {
		got := Median(test.array)
		assert.Equal(t, test.want, got)
	}
}

func TestMode(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  []int
	}{
		{
			name:  "array length == 0",
			array: []int{},
			want:  []int{},
		},
		{
			name:  "return one value",
			array: []int{1, 1, 2, 3, 5, 8, 13, 21, 34},
			want:  []int{1},
		},
		{
			name:  "return two value",
			array: []int{0, 0, 1, 1, 2, 3, 5, 8, 13, 21},
			want:  []int{0, 1},
		},
	}

	for _, test := range tests {
		got := Mode(test.array)
		assert.Equal(t, test.want, got)
	}

}
