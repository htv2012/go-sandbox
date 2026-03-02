// Demo: iota, which increases value
package main

import "fmt"

func main() {
	const (
		UserCatNormal = iota
		UserCatAdmin
		UserCatSuperUser
		UserCatService
	)

	fmt.Println("UserCatNormal", "=", UserCatNormal)
	fmt.Println("UserCatAdmin", "=", UserCatAdmin)
	fmt.Println("UserCatSuperUser", "=", UserCatSuperUser)
	fmt.Println("UserCatService", "=", UserCatService)
	fmt.Println()

	const (
		//_  = iota // skip 0
		B  = 1 << (10 * iota)
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
	)

	fmt.Println("1B = ", B, "bytes")
	fmt.Println("1KB = ", KB, "bytes")
	fmt.Println("1MB = ", MB, "bytes")
	fmt.Println("1GB = ", GB, "bytes")
	fmt.Println()
}
