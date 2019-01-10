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
}
