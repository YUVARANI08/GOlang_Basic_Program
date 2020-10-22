package main

import (
	"fmt"
)

func main() {
	n := []int{33, 44, 55, 66, 77}
	for i, value := range n {
		fmt.Println(value, i)
	}
	//slice the array
	slice := n[1:3]
	fmt.Println(slice)

	//copy the array
	newarray := make([]int, 5, 10)
	copy(newarray, n)
	fmt.Println(newarray)

	//append the value in array
	slice1 := append(newarray, 5, 10)
	fmt.Println(slice1)

}
