package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "[2606:4700:4700::1111]:53")
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	fmt.Printf("%T: %v\n", conn.LocalAddr(), conn.LocalAddr())
	fmt.Printf("%T: %v\n", conn.LocalAddr().Network(), conn.LocalAddr().Network())
	defer conn.Close()
}
