package main

import (
	"fmt"
)

func main() {

	var n [5]int
	n[0] = 10
	n[1] = 20
	n[2] = 30
	n[3] = 40
	n[4] = 50

	fmt.Println(n[0])

	//array usng looping
	for i, value := range n {
		fmt.Println(value, i)
	}

}
