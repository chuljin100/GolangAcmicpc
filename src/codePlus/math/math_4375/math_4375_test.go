package math4375_test

import (
	math4375 "codePlus/math/math_4375"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestMath4375(t *testing.T) {
	startTime := time.Now()
	log.Println("start")
	for i := 1; i < 10000; i++ {
		// if i%2 != 0 && i%5 != 0 {
		fmt.Println("i : ", i)
		// 	math4375.Calc(i)
		// }
		math4375.Calc(i)
	}
	elapsedTime := time.Since(startTime)

	fmt.Printf("실행시간: %f\n", elapsedTime.Seconds())
	// math4375.Calc(3)
	// math4375.Calc(2)
	// math4375.Calc(7)
	math4375.Calc(9901)
}
