package main

import (
	"fmt"
)

func PointerTest(temp *int) {
	*temp += 10
}

func main() {

	a := 0
	p := &a
	fmt.Println("修改前*p =", *p)
	PointerTest(p)
	fmt.Println("修改后*p =", *p)
}
