package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	l := lenStr(strs)
	var res string = ""
	flag := true
	for i := 0; i < l; i++ {
		temp := strs[0][i]

		for _, s := range strs {
			if s[i] != temp {
				flag = false
				break

			}
		}
		if flag {
			res += string(temp)
		}
	}

	fmt.Println(res)
	return res
}

func lenStr(strs []string) int {
	l := len(strs[0])
	for _, s := range strs {
		if len(s) < l {
			l = len(s)
		}
	}
	return l
}

func main() {
	var strs = []string{"cir", "car"}
	longestCommonPrefix(strs)
}
