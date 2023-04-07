package main

import (
	"answer/functions"
	"fmt"
)

func main() {
	var i int
	fmt.Print("select problem (1-3) : ")
	fmt.Scan(&i)

	if i == 1 {
		fmt.Print("\n1.interger to roman\ninput : ")
		fmt.Scan(&i)
		roman := functions.IntToRoman(i)
		fmt.Println(roman)
	}
	if i == 2 {
		fmt.Print("\n2.simple bank system")
		fmt.Scan(&i)

	}
}
