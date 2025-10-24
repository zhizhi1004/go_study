package main

import (
	"fmt"
)

const PI = 3.14

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

func (r Circle) Area() float64 {
	return PI * r.radius * r.radius
}

func (r Circle) Perimeter() float64 {
	return 2 * PI * r.radius

}

func printShapeInfo(s Shape) {
	fmt.Printf("面积 : %.2f\n", s.Area())
	fmt.Printf("周长 : %.2f\n", s.Perimeter())
}
func main() {

	fmt.Println("Rectangle info :")
	r := Rectangle{length: 20, breadth: 10}
	printShapeInfo(r)

	fmt.Println("Circle info :")
	c := Circle{radius: 5}
	printShapeInfo(c)
}
