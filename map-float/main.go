package main

import "fmt"

func main() {
	m := make(map[string]float64)
	m["a"] = 0
	fmt.Printf("%T", m["a"])
}
