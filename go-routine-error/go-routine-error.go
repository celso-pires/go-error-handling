// Write your answer here, and then test your code.
// Your job is to implement the findLargest() method.

package main

import (
	"fmt"
)

func isEven(n int) error {
	if n%2 == 0 {
		return fmt.Errorf("%d is an even number", n)
	}
	return nil
}

// isOdd returns a string channel containing
// whether a value is odd
func isOdd(n int) <-chan string {
	c := make(chan string)

	go func() {
		if err := isEven(n); err != nil {
			c <- fmt.Sprintf("Error: %v\n", err)
		}
		c <- fmt.Sprintf("%d is an odd number\n", n)
	}()
	return c
}

func main() {

	// This is how your code will be called.
	// Your answer should be the largest value in the numbers array.
	// You can edit this code to try different testing cases.
	var result string
	for i := 0; i < 255; i++ {
		result += <-isOdd(i)
		// time.Sleep(time.Millisecond)
		// runtime.Gosched()
	}

	fmt.Println(result)
}
