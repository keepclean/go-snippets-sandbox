package main

import "fmt"

var globalVar = "globalVar"

func main() {
	fmt.Println("before func: ", globalVar)
	doSomething(globalVar)
	fmt.Println("after func: ", globalVar)
}

func doSomething(s string) {
	fmt.Println("before closure: ", globalVar)
	globalVar = "localGlobalVar"
	func(x string) {
		fmt.Println("inside closure: ", x)
		callMe()
	}(globalVar)
	fmt.Println("after closure: ", globalVar)
}

func callMe() {
	fmt.Println("in callMe func: ", globalVar)
}
