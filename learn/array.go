// array.go
package main

import "fmt"

func main() {
	// Arrays are initialized to zeros
	var zeros [3]int
	fmt.Println("zeros =", zeros)
	fmt.Println()

	// Sparsed array
	var sparsed = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println("sparsed =", sparsed)
	sparsed[1] = 111
	fmt.Println("Set sparsed[1] = 111")
	fmt.Println("sparsed =", sparsed)
	fmt.Println("sparsed has", len(sparsed), "elements")
	fmt.Println()

	// Compare
	var ar1 = [...]int{1, 2, 3}
	var ar2 = [3]int{1, 2, 3}
	fmt.Println("ar1 =", ar1)
	fmt.Println("ar2 =", ar2)
	fmt.Println("Are they equal?", ar1 == ar2)
	fmt.Println()
}
