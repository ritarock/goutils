package array

func ToFloat64(arr interface{}) []float64 {
	var result []float64

	switch arr.(type) {
	case []float64:
		result = arr.([]float64)
	case []int:
		for _, v := range arr.([]int) {
			result = append(result, float64(v))
		}
	default:
		result = append(result, 0)
	}
	return result
}

func SumArray(arr interface{}) float64 {
	var sum float64 = 0
	newArr := ToFloat64(arr)

	for _, v := range newArr {
		sum += v
	}
	return sum
}

func MaxArray(arr interface{}) float64 {
	newArr := ToFloat64(arr)
	max := newArr[0]
	for _, v := range newArr {
		if max < v {
			max = v
		}
	}
	return max
}

func MinArray(arr interface{}) float64 {
	newArr := ToFloat64(arr)
	min := newArr[0]
	for _, v := range newArr {
		if v < min {
			min = v
		}
	}
	return min
}

func AvarageArray(arr interface{}) float64 {
	newArr := ToFloat64(arr)
	avarage := SumArray(newArr) / float64(len(newArr))
	return avarage
}
