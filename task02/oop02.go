package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (p Employee) printInfo() {
	fmt.Println("Name:", p.name, "Age:", p.age, "EmployeeID:", p.EmployeeID)
}

func main() {
	p1 := Person{"Mike", 30}
	e1 := Employee{p1, 1001}
	e1.printInfo()

}
