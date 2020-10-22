package main

import "fmt"

func main() {

	r1 := Student{"uva", 1234, 67, 89}
	r2 := Student{"viji", 234, 78, 90}
	fmt.Println(Calc(r1))
	fmt.Println(Calc(r2))
	fmt.Println(r1.name, r1.roll, "total:", Calc(r1))
	fmt.Println(r2.name, r2.roll, "total:", Calc(r2))

}

type Student struct {
	name string
	roll int
	m1   int
	m2   int
}

func Calc(r Student) int {

	return r.m1 + r.m2
}
