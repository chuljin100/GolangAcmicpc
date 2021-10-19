package math_9613_test

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	line := 3
	counterList := []int{4, 3, 3}
	valueList := [][]int{
		{10, 20, 30, 40},
		{7, 5, 12},
		{125, 15, 25},
	}

	for i := 0; i < line; i++ {
		counter := counterList[i]
		values := make([]int, counter)
		for j := 0; j < counter; j++ {
			values[j] = valueList[i][j]
		}
		result := 0

		for k := 0; k < counter; k++ {
			for l := k + 1; l < counter; l++ {
				result += GCD(values[k], values[l])
			}
		}
		fmt.Println(result)
	}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
