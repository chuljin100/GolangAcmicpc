package math_10430

import (
	"fmt"
)

//백준 답안 제출 코드
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	Calc(a, b, c)
}

func Calc(a int, b int, c int) {
	fmt.Println((a + b) % c)
	fmt.Println(((a % c) + (b % c)) % c)
	fmt.Println((a * b) % c)
	fmt.Println(((a % c) * (b % c)) % c)

}
