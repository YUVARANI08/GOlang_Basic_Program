package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("enter the table value")
	fmt.Scanln(&n)
	for i := 1; i <= 10; i++ {

		Multi(n)
		fmt.Println(n, "x", i, "=", n*i)
	}

}

func Multi(n int) int {
	return n

}
