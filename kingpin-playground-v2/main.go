package main

import (
	"fmt"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

var interval = kingpin.Flag("interval", "interval between sequential actions").Default("1s").Duration()

func init() {
	kingpin.Version("0.01")
	kingpin.Parse()
}

func main() {
	fmt.Printf("%T: %v\n", *interval, *interval)
	t1 := time.Now()
	// apparently this is wrong
	// tSleep := time.Second * time.Duration(*interval) // gives sleep for 277777h46m40s
	time.Sleep(*interval)
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}
