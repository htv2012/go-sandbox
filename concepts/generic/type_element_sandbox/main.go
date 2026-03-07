package main

import "fmt"
import "errors"
import "reflect"

type BigInt int64

type Integer interface {
	int | int8 | int16 | int32 | ~int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

func divMod[T Integer](numerator, denominator T) (T, T, error) {
	if denominator == 0 {
		return 0, 0, errors.New("divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func demo[T int | int8 | int16 | int32 | ~int64 | uint | uint8 | uint16 | uint32 | uint64](num, dem T) {
	a, b, err := divMod(num, dem)
	typeName := reflect.TypeOf(num).Kind()

	fmt.Printf("divMod[%s](%d, %d) -> ", typeName, num, dem)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d, %d\n", a, b)
	}
}

func main() {
	fmt.Println("# divMod Demo")
	demo(-15, -2)
	demo[uint](15, 2)
	demo[uint64](15, 0)

	// Demo: The use of ~int64
	var num1 BigInt = 15
	var dem1 BigInt = 2
	demo(num1, dem1)

	fmt.Println()
}
