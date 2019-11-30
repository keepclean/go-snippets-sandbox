package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type createUpdateOptions struct {
	CPU       float64
	Memory    int
	Instances int
}

func newCreateUpdateOptions(cmd *kingpin.CmdClause) *createUpdateOptions {
	options := &createUpdateOptions{}
	cmd.Flag("cpu", "cpu share").Short('c').Required().FloatVar(&options.CPU)
	cmd.Flag("mem", "mem share (integer MB)").Short('m').Required().IntVar(&options.Memory)
	if cmd.FullCommand() == "update" {
		cmd.Flag("instances", "instance count").Short('i').Required().IntVar(&options.Instances)
	}
	return options
}

var (
	app           = kingpin.New("myapp", "help me")
	create        = app.Command("create", "initial create/deploy of an app")
	createOptions = newCreateUpdateOptions(create)

	update        = app.Command("update", "update definition of an app (automatically deploys new definition)")
	updateOptions = newCreateUpdateOptions(update)

	urlCheck = app.Command("check-url", "URL for checking")
	urlFlag  = urlCheck.Flag("url", "URL").Default("http://example.com/").URL()
)

func main() {
	kingpin.Version("0.0.1")
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case urlCheck.FullCommand():
		fmt.Println(*urlFlag)
		path := &url.URL{Path: "/path"}
		u := *urlFlag
		fmt.Println(u.ResolveReference(path))
	}

	fmt.Fprintln(ioutil.Discard, createOptions, updateOptions)
}
