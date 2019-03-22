package main

import (
	"fmt"
	"math/rand"
	"time"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var interval = kingpin.Flag("interval", "interval between sequential actions").Default("5s").Duration()
var divisor = kingpin.Flag("divisor", "divisor for changing probablity of adding to slice").Default("3").Int()

func init() {
	kingpin.Version("0.01")
	kingpin.Parse()
}

func main() {
	tWin := *interval * 5
	fmt.Println("Time window:", tWin)
	// fmt.Printf("%T: %v\n", *interval, *interval)
	// apparently this is wrong
	// tSleep := time.Second * time.Duration(*interval) // gives sleep for 277777h46m40s
	ttl := []time.Time{}

	for {
		t1 := time.Now()

		rand.Seed(time.Now().UnixNano())
		r := rand.Int()
		if r%*divisor == 0 {
			fmt.Println("hit at", t1)
			ttl = append(ttl, t1)
		}

		firstTS := t1
		if len(ttl) > 0 {
			firstTS = ttl[0]
		}

		if t1.Sub(firstTS) > tWin {
			ttl = ttl[1:]
			fmt.Println("Removed element from ttl")
		}
		if len(ttl) >= 3 && t1.Sub(firstTS) < tWin {
			fmt.Println("len of ttl more than 3 in tWin:", ttl)
			break
		}

		fmt.Println("sleeping", *interval, ttl)
		time.Sleep(*interval)
	}
}
