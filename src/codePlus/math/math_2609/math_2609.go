package math_2609

import "fmt"

//백준 답안 제출 코드
func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(GCD(a, b))
	fmt.Println(LCM(a, b))
}
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}
