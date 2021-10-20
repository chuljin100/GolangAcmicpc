package math_1978

import (
	"fmt"
	"math"
)

func Math_1978() {
	var count int
	result := 0
	fmt.Scan(&count)

	for i := 0; i < count; i++ {
		var tmp int
		fmt.Scan(&tmp)
		if IsPrime(tmp) {
			result += 1
		}

	}
	fmt.Println(result)
}

func IsPrime(val int) bool {

	if val <= 1 {
		return false
	}

	if val == 2 {
		return true
	}

	if val%2 == 0 {
		return false
	}
	limit := int(math.Sqrt(float64(val)))
	fmt.Printf(" limit : %d", limit)

	for i := 2; i <= limit; i++ {
		if val%i == 0 {
			return false
		}
	}
	return true

}
