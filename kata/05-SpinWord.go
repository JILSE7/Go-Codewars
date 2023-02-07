package kata

import (
	"fmt"
	"strings"
)

/*
	Write a function that takes in a string of one or more words, and returns the same string, but with all five or more letter words reversed (Just like the name of this Kata). Strings passed in will consist of only letters and spaces. Spaces will be included only when more than one word is present.

	Examples:

	spinWords( "Hey fellow warriors" ) => returns "Hey wollef sroirraw"
	spinWords( "This is a test") => returns "This is a test"
	spinWords( "This is another test" )=> returns "This is rehtona test"
*/

func SpinWords(str string) string {

	strArray := strings.Split(str, " ")
	spin := []string{}
	for _, value := range strArray {
		p := strings.Split(value, "")
		if len(p) > 0 {
			if len(p) >= 5 {
				var newP string
				lengthP := len(p)
				fmt.Println(lengthP)
				for i := 1; i <= lengthP; i++ {
					newP += p[lengthP-i]
				}
				spin = append(spin, newP)
			} else {
				spin = append(spin, value)
			}

		}
	}

	return strings.Join(spin, " ")
} // SpinWords
