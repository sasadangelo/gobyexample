package main

import "fmt"

func main() {
	//========================================
	// Arithmetic operators
	//========================================
	// result is 10
	var myVariable1 int = 5 + 5
	fmt.Println(myVariable1)

	// result is 2.5
	var myVariable2 float32 = 5.0 / 2
	fmt.Printf("%.1f\n", myVariable2)

	// result is 10
	var myVariable3 int = 5 * 2
	fmt.Println(myVariable3)

	// result is 3
	var myVariable4 uint8 = 5 - 2
	fmt.Println(myVariable4)

	//========================================
	//Bitwise operators
	//========================================
	// result is 7
	var myVariable5 int = 5 | 2
	fmt.Println(myVariable5)

	// result is 2
	var myVariable6 float32 = 7 & 2
	fmt.Println(myVariable6)

	//========================================
	// Conditional operators
	//========================================
	// result is false
	var myVariable7 bool = true && false
	fmt.Println(myVariable7)

	// result is true
	var myVariable8 bool = true || false
	fmt.Println(myVariable8)

	// result is false
	var myVariable9 bool = !true
	fmt.Println(myVariable9)
}
