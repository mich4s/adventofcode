package main

import (
	"fmt"
	"strconv"
	"strings"
)

func createIntArray(input int) []int {
	response := make([]int, 0)
	numbers := strings.Split(fmt.Sprintln(input), "")
	for _, number := range numbers {
		intNumber, _ := strconv.Atoi(number)
		response = append(response, intNumber)
	}
	return response
}

func validateDuplicationCriteria(digits []int) bool {
	for index := range digits {
		if index < len(digits)-2 {
			nextValue := digits[index+1]
			thirdValue := digits[index+2]
			if digits[index] == nextValue && digits[index] != thirdValue {
				return true
			}
		}
	}
	return false
}

func validateDuplicationOnlyTwoCriteria(digits []int) bool {
	// This is lame as fck...
	return digits[0] == digits[1] && digits[1] != digits[2] ||
		digits[0] != digits[1] && digits[1] == digits[2] && digits[2] != digits[3] ||
		digits[1] != digits[2] && digits[2] == digits[3] && digits[3] != digits[4] ||
		digits[2] != digits[3] && digits[3] == digits[4] && digits[4] != digits[5] ||
		digits[3] != digits[4] && digits[4] == digits[5]
}

func validateDecreasingCriteria(digits []int) bool {
	for index := range digits {
		if index < len(digits)-2 {
			nextValue := digits[index+1]
			if nextValue < digits[index] {
				return false
			}
		}
	}
	return true
}

func validateSingleNumber(number int) bool {
	numbers := createIntArray(number)
	if !validateDecreasingCriteria(numbers) {
		return false
	}
	if !validateDuplicationCriteria(numbers) {
		return false
	}
	if !validateDuplicationOnlyTwoCriteria(numbers) {
		return false
	}
	return true
}

func calculate(from int, to int) int {
	response := 0
	for i := from; i <= to; i++ {
		if validateSingleNumber(i) {
			response++
		}
	}
	return response
}

func main() {
	response := calculate(278384, 824795)
	fmt.Println(response)
}
