// Two ways to create an error: errors.New() and fmt.Errorf
package main

import (
	"errors"
	"fmt"
)

func doubleOdd(n int) (int, error) {
	if n%2 == 0 {
		return 0, errors.New("expect odd number")
	}
	return n * 2, nil
}

func doubleDemo() {
	result, err := doubleOdd(12)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("double 12 ->", result)
}

func halfEven(n int) (int, error) {
	if n%2 == 1 {
		return 0, fmt.Errorf("expect even number, but got %d", n)
	}
	return n / 2, nil

}

func halfDemo() {
	result, err := halfEven(15)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("half 15 ->", result)
}
func main() {
	doubleDemo()
	halfDemo()
}
