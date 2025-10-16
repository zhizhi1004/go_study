package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	num := strconv.Itoa(x)
	for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 {
		if num[i] != num[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(121))

}
