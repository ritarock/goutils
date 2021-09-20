package array

import (
	"reflect"
	"testing"
)

func TestSortArray(t *testing.T) {
	type args struct {
		array interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "int array",
			args: args{array: []int{5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "float array",
			args: args{array: []float64{5, 4, 3, 2, 1}},
			want: []float64{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseSortArray(t *testing.T) {
	type args struct {
		array interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "int array",
			args: args{array: []int{1, 2, 3, 4, 5}},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "float array",
			args: args{array: []float64{1, 2, 3, 4, 5}},
			want: []float64{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseSortArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseSortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniueArray(t *testing.T) {
	type args struct {
		array interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "int array",
			args: args{array: []int{1, 2, 3, 3, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "float array",
			args: args{array: []float64{1, 2, 3, 3, 4, 5}},
			want: []float64{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniueArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniueArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxOfArray(t *testing.T) {
	type args struct {
		array interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "int array",
			args: args{array: []int{1, 2, 3, 4, 5}},
			want: 5,
		},
		{
			name: "float array",
			args: args{array: []float64{1, 2, 3, 4, 5}},
			want: 5.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxOfArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxOfArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinOfArray(t *testing.T) {
	type args struct {
		array interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "int array",
			args: args{array: []int{1, 2, 3, 4, 5}},
			want: 1,
		},
		{
			name: "float array",
			args: args{array: []float64{1, 2, 3, 4, 5}},
			want: 1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinOfArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinOfArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumOfArray(t *testing.T) {
	type args struct {
		array interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "int array",
			args: args{array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 55,
		},
		{
			name: "float array",
			args: args{array: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 55.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumOfArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAvarageOfArray(t *testing.T) {
	type args struct {
		array interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "int array",
			args: args{array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 5.5,
		},
		{
			name: "float array",
			args: args{array: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 5.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AvarageOfArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AvarageOfArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
