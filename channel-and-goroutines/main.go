package main

import (
	"fmt"
	"time"
)

func producer(start, end uint, chnl chan uint) {
	for i := start; i < end; i++ {
		chnl <- i
	}
}

func consumer(ch chan uint) {
	for {
		v, ok := <-ch
		if ok == false {
			fmt.Println("NOK")
			continue
		}
		fmt.Println("Received ", v, ok)
	}
}

func gen(ch chan uint) {
	data := map[uint]uint{
		0:   100,
		100: 200,
		200: 300,
		300: 400,
		400: 500,
	}

	for start, end := range data {
		producer(start, end, ch)
	}
}

func main() {
	ch := make(chan uint)
	go consumer(ch)
	gen(ch)

	time.Sleep(100 * time.Second)
}
