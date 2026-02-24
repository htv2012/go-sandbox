package main

import "fmt"

func banner(text string) {
	fmt.Println("\n#", text)
}

func note(text string) {
	fmt.Println("#", text)
}

func show[L any, T any](label L, value T) {
	fmt.Printf("  %v: %v\n", label, value)
}

func nil_map_demo() {
	var nilMap map[string]int
	banner("Nil Map Demo")
	show("nilMap", nilMap)
	show("nilMap == nil?", nilMap == nil)
	show("nilMap[foo]", nilMap["foo"])
}

func basic_demo() {
	score := map[string]int{
		"Anna": 89,
		"Jake": 87,
	}
	banner("Basic Demo")
	show("score", score)
}

func make_demo() {
	ppm := make(map[string]int, 3)
	banner("Make a map with 3 slots")
	ppm["Peter"] = 1
	ppm["Paul"] = 2
	ppm["Mary"] = 3
	show("ppm", ppm)

	note("Add another key/value")
	ppm["John"] = 4
	show("ppm", ppm)
}

func counter_demo() {
	banner("Counter demo")
	counter := map[string]int{}

	counter["foo"] = 3
	counter["bar"]++ // default=0, after increment=1
	show("counter", counter)

}
func main() {
	nil_map_demo()
	basic_demo()
	make_demo()
	counter_demo()
}
