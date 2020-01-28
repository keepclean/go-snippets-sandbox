package main

import (
	"fmt"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cwd)
}
