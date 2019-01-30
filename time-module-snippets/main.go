package main

import (
	"fmt"
	"net/url"
	"time"
)

func main() {
	t, _ := time.Parse(time.RFC3339, "2018-12-20T20:00:00Z")
	fmt.Println(t)

	got := t.Format("2006-01-02 15:04")
	fmt.Println(got)
	fmt.Println(t.Weekday())

	var startdate string
	vals := make(url.Values)
	vals.Set("include_oncall", "true")
	vals.Set("since", startdate)

	fmt.Println(vals.Encode())

	t1 := time.Now()
	fmt.Println(t1.AddDate(0, 0, 7))
	fmt.Println(t1.AddDate(0, 0, 7).Format("2006-01-02"))
	if _, err := time.Parse("2006-01-02", "2019-01-37"); err != nil {
		fmt.Println(err)
	}

}
