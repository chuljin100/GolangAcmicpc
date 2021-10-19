package math_9613

import "fmt"

func main() {
	var line, counter int
	fmt.Scan(&line)
	for i := 0; i < line; i++ {
		fmt.Scan(&counter)
		var tmp int
		values := make([]int, counter)
		for j := 0; j < counter; j++ {
			fmt.Scan(&tmp)
			values[j] = tmp
			tmp = 0
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
