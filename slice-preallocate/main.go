package main

import "fmt"

func main() {
	dict := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	keys := make([]string, 0, len(dict))

	for key := range dict {
		keys = append(keys, key)
	}

	fmt.Println(keys)
}
