package main

import (
	"fmt"
)

func PointerAndSlice(p *[]int) {
	for i := 0; i < len(*p); i++ {
		(*p)[i] = (*p)[i] * 2
	}
}
func main() {
	var temp []int = []int{1, 2, 3, 4, 5}
	for _, value := range temp {
		fmt.Println("修改前value = ", value)
	}

	PointerAndSlice(&temp)

	for _, value := range temp {
		fmt.Println("修改后value = ", value)
	}

}
