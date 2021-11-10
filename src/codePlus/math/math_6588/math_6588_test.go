package math_6588_test

import (
	a "codePlus/math/math_6588"
	"fmt"
	"testing"
	"time"
)

func TestMath6588(t *testing.T) {
	startTime := time.Now()

	primeList := a.MakePrimeNumberList(1000000)
	// for i := 4; i < 1000000; i += 2 {
	// 	a.CheckPrime(primeList, i)
	// }
	a.CheckPrime(primeList, 8)
	a.CheckPrime(primeList, 20)
	a.CheckPrime(primeList, 42)
	elapsedTime := time.Since(startTime)

	fmt.Printf("실행시간: %s\n", elapsedTime)
}

func TestPrimeNumberList(t *testing.T) {
	startTime := time.Now()
	list := a.MakePrimeNumberList(1000000)
	fmt.Println(list)
	for i := 4; i < 1000000; i += 2 {
		a.CheckPrime(list, i)
	}
	elapsedTime := time.Since(startTime)

	fmt.Printf("실행시간: %s\n", elapsedTime)
}
func TestBool(t *testing.T) {
	testList := make([]bool, 100)
	for i := 0; i < len(testList); i++ {
		fmt.Println(testList[i])
	}
}
