package math_1934_test

import (
	a "codePlus/math/math_1934"
	"fmt"
	"testing"
)

func TestT(t *testing.T) {
	count := 3
	input := [][]int{{1, 45000}, {6, 10}, {13, 17}}

	for i := 0; i < count; i++ {
		fmt.Println(a.LCM(input[i][0], input[i][1]))
	}
}
