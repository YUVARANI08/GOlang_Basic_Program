//FUNCTIONS
package main

import "fmt"

func calc(n1, n2 int) (int, int) {
	o1 := n1 + n2
	o2 := n1 - n2
	return o1, o2
}
func main() {
	num1 := 4
	num2 := 5
	result, result1 := calc(num1, num2)
	fmt.Println(result, result1)
	var m = 8
	display(m)
}

//odd even
func display(n int) {
	var m = n%2 == 0
	fmt.Println("even", m)
	var m1 = n%2 != 0
	fmt.Println("odd", m1)

}
