package math4375

import (
	"fmt"
	"time"
)

func Math_4375() {
	for {
		var tmp int
		fmt.Scan(&tmp)
		Calc(tmp)
	}
}

func Calc(input int) {
	startTime := time.Now()
	defer func() {
		if err := recover(); err != nil {

		}
	}()
	if input%2 == 0 || input%5 == 0 {
		panic(1)
	}

	cnt := "1"
	num := 1
	for {
		if (num % input) == 0 {
			fmt.Println(len(cnt))
			break
		}
		num %= input
		num = (num * 10) + 1
		cnt += "1"
	}
	elapsedTime := time.Since(startTime)

	fmt.Printf("실행시간: %f\n", elapsedTime.Seconds())

}
