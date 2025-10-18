package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	digits = append(digits, 0)
	digits[0] = 1
	return digits
}

func main() {
	digits := []int{1, 2, 3, 9}
	fmt.Println(plusOne(digits))
}
