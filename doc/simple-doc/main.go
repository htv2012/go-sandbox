// Package main drives the whole project.
package main

import "fmt"

// Money represents the amount of money (e.g. 1.75) and the currency it
// is in (e.g. "USD", "CDN", "VND"...)
type Money struct {
	Value    float64
	Currency string
}

func (m Money) String() string {
	return fmt.Sprintf("%.2f%s", m.Value, m.Currency)
}

/*
Convert converts amount of currency from src to dest.

Currently, we can only convert USD -> CDN
*/
func Convert(src Money, dest string) (Money, error) {
	// Currently can only convert USD to CDN
	if src.Currency == "USD" && dest == "CDN" {
		return Money{Value: src.Value * 1.37, Currency: dest}, nil
	}

	return Money{}, fmt.Errorf("Cannot convert %v to %s", src, dest)
}

// main drives the whole project. This is an entry point.
func main() {
	src := Money{Value: 100.0, Currency: "USD"}
	dest, err := Convert(src, "CDN")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v = %v\n", src, dest)
	}

	dest, err = Convert(src, "VND")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v = %v\n", src, dest)
	}
}
