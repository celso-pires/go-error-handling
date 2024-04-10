// Write your answer here, and then test your code.
// Your job is to implement the findLargest() method.

package main

import (
	"fmt"
)

// isOdd() returns a panic error message if the number is not even
func isOdd(n int) {
	panic(fmt.Sprintf("%d is not even", n))
}

// is even returns a boolean and prints a recover message
func isEven(n int) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered: %v\n", r)
		}
	}()

	if n%2 != 0 {
		isOdd(n)
	}
	return true
}

func main() {

	// This is how your code will be called.
	// You can edit this code to try different testing cases.

	var answer string
	for i := 0; i < 5; i++ {
		result := isEven(i)
		answer += fmt.Sprintf("isEven() for number %d, returns %t\n", i, result)
	}
	fmt.Println(answer)
}
