package main

import "fmt"

func doMath(num, dem int) {
	fmt.Printf("%v/%v = ", num, dem)
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("PANIC:", msg)
		}
	}()

	fmt.Println(num / dem)
}

func main() {
	for _, dem := range []int{2.0, 3.0, 0.0, 4.0} {
		doMath(100.0, dem)
	}
}
