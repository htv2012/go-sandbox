// Demo: iota, which increases value
package main

import "fmt"

const (
	B  = 1 << (10 * iota)
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
)

func ConvertStorage(value float64, fromUnit int, toUnit int) float64 {
	value = value * float64(fromUnit) / float64(toUnit)
	return value
}

func main() {
	fmt.Println("1B = ", B, "bytes")
	fmt.Println("1KB = ", KB, "bytes")
	fmt.Println("1MB = ", MB, "bytes")
	fmt.Println("1GB = ", GB, "bytes")
	fmt.Println()

	type TestCase struct {
		Value     float64
		FromUnit  int
		FromLabel string
		ToUnit    int
		ToLabel   string
	}

	testCases := []TestCase{
		{Value: 1.5, FromUnit: TB, FromLabel: "TB", ToUnit: GB, ToLabel: "GB"},
		{Value: 2, FromUnit: KB, FromLabel: "KB", ToUnit: B, ToLabel: "B"},
		{Value: 2, FromUnit: MB, FromLabel: "MB", ToUnit: B, ToLabel: "B"},
	}

	for _, tc := range testCases {
		converted := ConvertStorage(tc.Value, tc.FromUnit, tc.ToUnit)
		fmt.Printf("%f %s -> %f %s\n", tc.Value, tc.FromLabel, converted, tc.ToLabel)
	}
}
