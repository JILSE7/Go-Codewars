package kata

import (
	"fmt"
	"strings"
)

/*
	The goal of this exercise is to convert a string to a new string where each character in the new string is "(" if that character appears only once in the original string, or ")"
	if that character appears more than once in the original string. Ignore capitalization when determining if a character is a duplicate.

	Examples
	"din"      =>  "((("
	"recede"   =>  "()()()"
	"Success"  =>  ")())())"
	"(( @"     =>  "))(("
	Notes
	Assertion messages may be unclear about what they display in some languages. If you read "...It Should encode XXX", the "XXX" is the expected result, not the input!
*/

func DuplicateEncode(word string) string {
	letters := make(map[string][]int)
	s := strings.Split(word, "")
	title := make([]string, len(s))

	for i, val1 := range s {
		letter := strings.ToLower(val1)
		letters[letter] = append(letters[val1], i)
	}

	for index, _ := range letters {
		letterArr := letters[index]

		if len(letterArr) > 1 {
			// fmt.Println("tengo que poner ) en ", index)
			for _, v := range letterArr {
				fmt.Println(v)
				title[v] = ")"
			}
		}

		if len(letterArr) == 1 {
			// fmt.Println(letterArr)
			title[letterArr[0]] = "("
			// fmt.Println("tengo que poner ( en", index)
		}
	}
	// fmt.Println(letters)
	return strings.Join(title, "")
}

/* func DuplicateEncode(word string) string {
	newString := ""
	for _, ch := range word {
			if strings.Count(strings.ToLower(word), strings.ToLower(string(ch))) > 1 {
					newString += ")"
			} else {
					newString += "("
			}
	}

	return newString
} */
