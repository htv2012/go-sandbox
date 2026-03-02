package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var stack Stack

	for {
		var line string
		fmt.Print("& ")
		if scanner.Scan() {
			line = scanner.Text()
		}
		if line == "q" {
			break
		}
		stack = EnstackLine(stack, line)
		fmt.Println(stack)
		tos, err := stack.Pop()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if tos == "+" {
			b, err := stack.PopFloat64()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("b =", b)
			a, err := stack.PopFloat64()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("a =", a)
			res := a + b
			fmt.Println("->", res)
			stack.Push(fmt.Sprintf("%f", res))
		}
		fmt.Println(stack)
	}
}
