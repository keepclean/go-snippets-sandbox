package main

import (
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
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
