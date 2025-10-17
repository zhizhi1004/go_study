package main

import "fmt"

func isValid(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}

	var strMaP = map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	var strSlince = []byte{}
	for i := 0; i < n; i++ {
		if strMaP[s[i]] > 0 {
			if len(strSlince) == 0 || strMaP[s[i]] != strSlince[len(strSlince)-1] {
				return false
			}
			strSlince = strSlince[:len(strSlince)-1]
		} else {
			strSlince = append(strSlince, s[i])
		}
	}
	return len(strSlince) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
}
