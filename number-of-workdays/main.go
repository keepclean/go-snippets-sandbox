package main

import (
	"fmt"
	"time"

	"github.com/rickar/cal"
)

func main() {
	t1, _ := time.Parse("2006-01-02", "2019-04-15")
	t2, _ := time.Parse("2006-01-02", "2019-04-29")

	cUK := cal.NewCalendar()
	cal.AddBritishHolidays(cUK)

	var n uint
	for i := 0; !t2.Equal(t1); i++ {
		if t1.Weekday() != 0 && t1.Weekday() != 6 && !cUK.IsHoliday(t1) {
			n++
		}
		t1 = t1.AddDate(0, 0, 1)
	}

	fmt.Println(n)
}
