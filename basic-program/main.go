package main

import "fmt"

func main() {
	// Variables
	var x int = 10
	y := 20

	// Control Structures
	fmt.Println(" If Condition:")
	if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("x is less than or equal to y")
	}

	fmt.Println("\n For Loop:")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Function call
	fmt.Println("\n Function:")
	result := add(x, y)
	fmt.Println("Result of addition:", result)
}

func add(a int, b int) int {
	return a + b
}
