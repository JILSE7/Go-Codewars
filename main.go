package main

import (
	"fmt"
	kata "github/JILSE7/codewars/kata"
)

func main() {

	// var a = kata.ToJadenCase("How can mirrors be real if our eyes aren't real")
	// a := kata.DuplicateEncode("(( @")

	//? TWO SUM
	/* test := []int{2, 2, 3}
	a := kata.TwoSum(test, 4)
	fmt.Println(a) */

	//? Persistence
	/* persistence := kata.Persistence(25)
	fmt.Println(persistence) */

	//? SpinWorld (palabra alreves)
	/* 	spinWorld := kata.SpinWords("Hey fellow warriors")
	   	fmt.Println(spinWorld) */

	//? RGB To Hex Conversion (255,255,255) -> #FFFFFF
	/* myHex := kata.RGB(-20, 275, 125)
	fmt.Printf("My ex %v", myHex) */

	//? Playing with passphrases
	// kata.PlayPass("3SI#!NLEQCB7   BEGXWIOKFD1R5H9VYC5N9 Y 7!A8ZP#N", 5)

	//? Is Prime
	isPrime := kata.IsPrime(3)

	fmt.Println(isPrime)
}
