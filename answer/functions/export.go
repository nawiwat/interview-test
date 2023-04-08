package functions

import (
	"fmt"
	"strconv"
	"strings"
)

//1.interger to roman

func IntToRoman(num int) string {

	//roman digit
	explanation := [13]struct {
		digit []string
		value int
	}{}
	digitValue := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	numConv, numExpl := num, num

	//convert to roman
	roman := []string{}
	for _, currentDigit := range digitValue {
		for numConv >= currentDigit.value {
			roman = append(roman, currentDigit.digit)
			numConv -= currentDigit.value
		}
	}
	answer := strings.Join(roman, "")
	fmt.Println("Output: \"" + answer + "\"")

	// explanation
    sample := ""
	var explanationStr []string
	var explanationAns []string
	for i, currentDigit := range digitValue {
		for numExpl >= currentDigit.value {
			explanation[i].digit = append(explanation[i].digit, currentDigit.digit)
			explanation[i].value += currentDigit.value
            sample = currentDigit.digit
			numExpl -= currentDigit.value
		}
	}
	for _, currentDigit := range explanation {
		for currentDigit.value != 0 {
			explanationStr = append(explanationStr, strings.Join(currentDigit.digit, ""))
			explanationStr = append(explanationStr, " = ")
			explanationStr = append(explanationStr, strconv.Itoa(currentDigit.value))
			exp := strings.Join(explanationStr, "")
			explanationAns = append(explanationAns, exp)
			explanationStr = explanationStr[:0]
            break
		}
	}
    if len(explanationAns) == 1 {
        unit := num
        fmt.Print("Explanation: ")
        if sample == "M" {
            unit = num / 1000
            fmt.Println(strconv.Itoa(num) + " is represented as " + strconv.Itoa(unit) + " thousands.")
        } else if sample == "C" {
            unit = num / 100
            fmt.Println(strconv.Itoa(num) + " is represented as " + strconv.Itoa(unit) + " hundreds.")
        } else if sample == "X" {
            unit = num / 10
            fmt.Println(strconv.Itoa(num) + " is represented as " + strconv.Itoa(unit) + " tens.")
        } else if sample == "I" {
            fmt.Println(strconv.Itoa(num) + " is represented as " + strconv.Itoa(unit) + " ones.")
        }
        
    } else {
        noAnd := make([]string, len(explanationAns)-1)
        copy(noAnd, explanationAns)
        answer = strings.Join(noAnd, ", ")
	    fmt.Println("Explanation: " + answer +" and "+ explanationAns[len(explanationAns)-1] + ".")
    }
	return ""
}

//3 two sum

func TwoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, num := range nums {
        a := target - num
        if j, ok := m[a]; ok {
            return []int{j, i}
        }
        m[num] = i
    }
    return nil
}

//5 array rotation
func Solution(A []int, K int) []int {
    N := len(A)
    if N == 0 {
        return A
    }
    K = K % N
    if K == 0 {
        return A
    }
    B := make([]int, N)
    for i := 0; i < N; i++ {
        B[(i+K)%N] = A[i]
    }
    return B
}