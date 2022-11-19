package main

import (
	"fmt"
)

func FunctionGolang() {
	var shapes []shape
	s := &square{side: 2}
	shapes = append(shapes, s)
	r := &rectangle{length: 2, breath: 3}
	shapes = append(shapes, r)
	for _, shape := range shapes {
		fmt.Printf("Type: %s, Area %d\n", shape.getType(), shape.area())
	}

}

type shape interface {
	area() int
	getType() string
}

type rectangle struct {
	length int
	breath int
}

func (r *rectangle) area() int {
	return r.length * r.breath
}

func (r *rectangle) getType() string {
	return "rectangle"
}

type square struct {
	side int
}

func (s *square) area() int {
	return s.side * s.side
}

func (s *square) getType() string {
	return "square"
}

type area func(int, int) int

func Print(x, y int, a area) {
	fmt.Printf("Area is : %d\n", a(x, y))
}

func getAreaFunc() area {
	return func(x, y int) int {
		return x * y
	}
}

func UserDefinedType() {
	areaF := getAreaFunc()
	Print(3, 4, areaF)
}

var Max = func(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func AnonymousFunc() {
	res := Max(2, 3)
	fmt.Println(res)
}

func InvokedFunction() {
	squareOf2 := func() int {
		return 2 * 2
	}()

	fmt.Println(squareOf2)
}

func AddVariadicFunc(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func VariadicFunc() {
	fmt.Println(AddVariadicFunc(1, 2))
	fmt.Println(AddVariadicFunc(1, 2, 3))
	fmt.Println(AddVariadicFunc(1, 2, 3, 4))
}
