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

func validateSingleNumber(number int) bool {
	numbers := createIntArray(number)
	flag := false
	for index := range numbers {
		if index < len(numbers)-2 {
			nextValue := numbers[index+1]
			if numbers[index] == nextValue {
				flag = true
			}
		}
	}
	if !flag {
		return false
	}
	for index := range numbers {
		if index < len(numbers)-2 {
			nextValue := numbers[index+1]
			if nextValue < numbers[index] {
				flag = false
				break
			}
		}
	}
	return flag
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
	// response := calculate(278384, 300084)
	// response := calculate(299999, 300000)
	fmt.Println(response)
}
