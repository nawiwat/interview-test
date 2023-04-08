package main

import (
	"bufio"
	"answer/functions"
	"fmt"
	"strings"
	"os"
)

func main() {
	var i int
	fmt.Print("select problem (1/3/5) : ")
	fmt.Scan(&i)
	//1 interger to roman
	if i == 1 {

		fmt.Print("1.interger to roman\ninput : ")
		fmt.Scan(&i)
		roman := functions.IntToRoman(i)
		fmt.Println(roman)
	}
	//3 two sum
	if i == 3 {
		fmt.Print("3.two sum\n")
		var target int
  		fmt.Print("Input: ")
  		var input string
  		fmt.Scan(&input)
		input = strings.Trim(input, "[]")
		var arr []int
		for _, s := range strings.Split(input, ",") {
			var n int
			fmt.Sscanf(s, "%d", &n)
			arr = append(arr, n)
		}
		fmt.Print("Target: ")
		fmt.Scan(&target)
		ans := functions.TwoSum(arr,target)
		if ans != nil {
			fmt.Print("Output: [")
			for i := 0; i < len(ans); i++ {
				fmt.Printf("%d", ans[i])
				if i < len(ans)-1 {
					fmt.Printf(",")
				}
			}
			fmt.Print("]")
		} else {
			fmt.Print("Output: no answer")
		}
	}

	//5 array rotation
	if i == 5 {
		fmt.Print("5.array rotation\n")
		fmt.Print("Input A : ")
		var K int
  		var input string
		fmt.Scanln(&input)
		reader := bufio.NewReader(os.Stdin)
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.Trim(input, "[]")
		var A []int
		for _, s := range strings.Split(input, ",") {
			var n int
			fmt.Sscanf(s, "%d", &n)
			A = append(A, n)
		}

		fmt.Print("Input K : ")
		fmt.Scan(&K)

		ans := functions.Solution(A,K)
		fmt.Print("Output: ",ans)
	}
}
