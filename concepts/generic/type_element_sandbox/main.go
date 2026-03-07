package main

import "fmt"
import "errors"

type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

func divMod[T Integer](numerator, denominator T) (T, T, error) {
	if denominator == 0 {
		return 0, 0, errors.New("divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func main() {
	fmt.Println("divMod Demo")

	tests := [][]uint{
		{15, 5},
		{15, 2},
		{15, 0},
	}

	for _, pair := range tests {
		num := pair[0]
		dem := pair[1]
		a, b, err := divMod(num, dem)
		fmt.Printf("  divMod(%d, %d) -> ", num, dem)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%d, %d\n", a, b)
		}
	}
}
