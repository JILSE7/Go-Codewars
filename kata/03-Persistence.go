package kata

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	 Write a function, persistence, that takes in a positive parameter num and returns its multiplicative persistence, which is the number of times you must multiply the digits in num until you reach a single digit.

	For example (Input --> Output):

	39 --> 3 (because 3*9 = 27, 2*7 = 14, 1*4 = 4 and 4 has only one digit)
	999 --> 4 (because 9*9*9 = 729, 7*2*9 = 126, 1*2*6 = 12, and finally 1*2 = 2)
	4 --> 0 (because 4 is already a one-digit number)
*/

func Persistence(n int) int {

	if n <= 9 {
		return 0
	}

	// aux := 1
	acc := 1
	op := 0
	fmt.Println("este es n", n)
	for n >= 10 {
		s2 := strconv.Itoa(n)
		fmt.Println("este es s2", s2)
		numbersArr := strings.Split(s2, "")
		fmt.Println("este es el arreglo", numbersArr)
		for i, _ := range numbersArr {
			num, _ := strconv.Atoi(numbersArr[i])
			acc = acc * num
		}

		fmt.Println("este es el acumulador", acc)

		n = acc
		op++
		acc = 1

	}

	fmt.Println(op)
	return op
}
