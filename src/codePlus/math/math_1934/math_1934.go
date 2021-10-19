package math_1934

import "fmt"

func main() {
	var counter, a, b int
	fmt.Scan(&counter)
	for i := 0; i < counter; i++ {
		fmt.Scan(&a, &b)
		fmt.Println(LCM(a, b))
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

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
