package main

import (
	"fmt"
	"reflect"
)

func main() {
	type BackendP struct {
		Host            string
		ConnectTimeout  *uint
		ResponseTimeout *uint
		Enabled         *bool
	}

	b1 := BackendP{
		Host: "backend01",
	}

	fmt.Printf("B1 ConnectTimeout Pointer: %v\n", &b1.ConnectTimeout)
	printStruct(b1)

	type Backend struct {
		Host            string
		ConnectTimeout  uint
		ResponseTimeout uint
		Enabled         bool
	}

	b2 := Backend{
		Host: "backend02",
	}

	fmt.Printf("B2 ConnectTimeout Pointer: %v\n", &b2.ConnectTimeout)
	printStruct(b2)
}

func printStruct(s interface{}) {
	fields := reflect.TypeOf(s)
	values := reflect.ValueOf(s)

	num := fields.NumField()
	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		fmt.Print(
			"Type: ", field.Type, "\n",
			"Name: ", field.Name, "\n",
			"Value: ", value, "\n",
			"CanAddr: ", value.CanAddr(), "\n",
			"-----------", "\n",
		)
	}
	fmt.Println("*********")
}
