package kata

import (
	"fmt"
	"strconv"
	"strings"
)

var abecedary []string = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func PlayPass(s string, n int) string {
	fmt.Println(s)
	strArray := strings.Split(s, "")
	fmt.Println(strArray, len(strArray))
	var newSentence string

	for _, letter := range strArray {
		isLetter := false
		upperLetter := strings.ToUpper(letter)

		for j, abc := range abecedary {
			shift := j + n
			if shift > len(abecedary)-1 {
				shift = shift - len(abecedary)
			}
			fmt.Println(shift)
			if abc == upperLetter {
				isLetter = true
				newSentence += abecedary[shift]
			}
		}

		if !isLetter {
			// si es un numero
			if number, err := strconv.Atoi(letter); err == nil {
				complement9 := 9 - number
				newSentence += fmt.Sprintf("%v", complement9)
			} else { // si es un caracter
				newSentence += letter
			}
		}

	}

	newStr := strings.Split(newSentence, "")
	var reverse string
	fmt.Println(newStr)
	for i := len(newStr) - 1; i >= 0; i-- {
		isLetter := false
		for _, abc := range abecedary {
			if abc == strings.ToUpper(newStr[i]) {
				isLetter = true
				if i%2 == 0 {
					reverse += newStr[i]
				} else {
					reverse += strings.ToLower(newStr[i])
				}
			}
		}

		if !isLetter {
			reverse += newStr[i]
		}
	}

	fmt.Println(newSentence, reverse)
	return reverse
}
