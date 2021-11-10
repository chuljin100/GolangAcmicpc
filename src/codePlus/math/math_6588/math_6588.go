package math_6588

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func Math_6588() {

	reader := bufio.NewReader(os.Stdin)

	primeList := MakePrimeNumberList(1000000)
	for {
		var tmp int
		fmt.Fscan(reader, &tmp)
		if tmp == 0 {
			break
		}
		CheckPrime(primeList, tmp)
	}
}
func CheckPrime(primeList []bool, checkNum int) {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
Calc:
	for i := 3; i <= checkNum/2; i++ {
		if !primeList[checkNum-i] && !primeList[i] {

			result := fmt.Sprintf("%d = %d + %d", checkNum, i, checkNum-i)
			fmt.Println(result)
			break Calc
		}
		if i == checkNum {
			fmt.Fprintln(writer, "Goldbach's conjecture is wrong.")
			break Calc
		}

	}
}

func MakePrimeNumberList(end int) []bool {

	primes := make([]bool, end)
	primes[0] = true
	primes[1] = true
	for i := 2; i < end; i++ {
		if IsPrime(i) {
			for j := i * i; j < end; j += i {
				primes[j] = true
			}
		}
	}

	return primes
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
	for i := 2; i <= limit; i++ {
		if val%i == 0 {
			return false
		}
	}
	return true

}
