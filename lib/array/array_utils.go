package array

import (
	"sort"
)

func SortArray(array interface{}) interface{} {
	switch array.(type) {
	case []int:
		sort.Sort(sort.IntSlice(array.([]int)))
	case []float64:
		sort.Sort(sort.Float64Slice(array.([]float64)))
	}

	return array
}

func ReverseSortArray(array interface{}) interface{} {
	switch array.(type) {
	case []int:
		sort.Sort(sort.Reverse(sort.IntSlice(array.([]int))))
	case []float64:
		sort.Sort(sort.Reverse(sort.Float64Slice(array.([]float64))))
	}
	return array
}

func UniqueArray(array interface{}) interface{} {
	var newList interface{}

	switch array.(type) {
	case []int:
		m := make(map[int]struct{})
		newList = make([]int, 0)

		for _, v := range array.([]int) {
			if _, ok := m[v]; !ok {
				m[v] = struct{}{}
				newList = append(newList.([]int), v)
			}
		}
	case []float64:
		m := make(map[float64]struct{})
		newList = make([]float64, 0)

		for _, v := range array.([]float64) {
			if _, ok := m[v]; !ok {
				m[v] = struct{}{}
				newList = append(newList.([]float64), v)
			}
		}
	}

	return newList
}

func MaxOfArray(array interface{}) interface{} {
	var max interface{}
	arr := ReverseSortArray(array)

	switch array.(type) {
	case []int:
		max = arr.([]int)[0]
	case []float64:
		max = arr.([]float64)[0]
	}

	return max
}

func MinOfArray(array interface{}) interface{} {
	var min interface{}
	arr := SortArray(array)

	switch array.(type) {
	case []int:
		min = arr.([]int)[0]
	case []float64:
		min = arr.([]float64)[0]
	}

	return min
}

func SumOfArray(array interface{}) interface{} {
	var sum interface{}

	switch array.(type) {
	case []int:
		var s int
		for _, v := range array.([]int) {
			s += v
		}
		sum = s
	case []float64:
		var s float64
		for _, v := range array.([]float64) {
			s += v
		}
		sum = s
	}

	return sum
}

func AvarageOfArray(array interface{}) interface{} {
	var avarage interface{}

	sumOfArray := SumOfArray(array)

	switch array.(type) {
	case []int:
		avarage = sumOfArray.(int) / len(array.([]int))
	case []float64:
		avarage = sumOfArray.(float64) / float64(len(array.([]float64)))
	}

	return avarage
}

func UnionArray(array ...interface{}) interface{} {
	var unionArray interface{}

	switch array[0].(type) {
	case []int:
		tmp := []int{}
		for _, v := range array {
			for _, v2 := range v.([]int) {
				tmp = append(tmp, v2)
			}
		}
		unionArray = UniqueArray(tmp)
	case []float64:
		tmp := []float64{}
		for _, v := range array {
			for _, v2 := range v.([]float64) {
				tmp = append(tmp, v2)
			}
		}
		unionArray = UniqueArray(tmp)
	}

	return unionArray
}

func IntersectArray(array ...interface{}) interface{} {
	var intersectArray interface{}

	switch array[0].(type) {
	case []int:
		tmp := []int{}
		tmp2 := []int{}
		if len(array) == 2 {
			for _, v := range array[0].([]int) {
				for _, v2 := range array[1].([]int) {
					if v == v2 {
						tmp2 = append(tmp2, v)
					}
				}
			}
			intersectArray = tmp2
		} else {
			for i := 0; i < len(array)-2; i++ {
				func(array1, array2 []int) {
					for _, v := range array1 {
						for _, v2 := range array2 {
							if v == v2 {
								tmp = append(tmp, v)
							}
						}
					}
				}(array[0].([]int), array[1].([]int))
				func(array1, array2 []int) {
					for _, v := range array1 {
						for _, v2 := range array2 {
							if v == v2 {
								tmp2 = append(tmp2, v)
							}
						}
					}
				}(tmp, array[i+2].([]int))
				intersectArray = tmp2
			}
		}

	case []float64:
		tmp := []float64{}
		tmp2 := []float64{}
		if len(array) == 2 {
			for _, v := range array[0].([]float64) {
				for _, v2 := range array[1].([]float64) {
					if v == v2 {
						tmp2 = append(tmp2, v)
					}
				}
			}
			intersectArray = tmp2
		} else {
			for i := 0; i < len(array)-2; i++ {
				func(array1, array2 []float64) {
					for _, v := range array1 {
						for _, v2 := range array2 {
							if v == v2 {
								tmp = append(tmp, v)
							}
						}
					}
				}(array[0].([]float64), array[1].([]float64))
				func(array1, array2 []float64) {
					for _, v := range array1 {
						for _, v2 := range array2 {
							if v == v2 {
								tmp2 = append(tmp2, v)
							}
						}
					}
				}(tmp, array[i+2].([]float64))
				intersectArray = tmp2
			}
		}
	}

	return intersectArray
}
