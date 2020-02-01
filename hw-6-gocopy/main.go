package main

import (
	"flag"
	"github.com/gookit/color"
	"os"
)

func main() {
	var from, to string
	var limit, offset int64
	var progressBarEnabled bool

	flag.StringVar(&from, "from", "", "/path/to/file")
	flag.StringVar(&to, "to", "", "/path/to/file")
	flag.Int64Var(&limit, "limit", 0, "1024")
	flag.Int64Var(&offset, "offset", 0, "1024")
	flag.BoolVar(&progressBarEnabled, "progress", true, "true/false")

	flag.Parse()

	c := NewGoCopy()

	err := c.Copy(from, to, limit, offset, progressBarEnabled)

	if err != nil {
		color.SetOutput(os.Stderr)
		color.Error.Println(err.Error())
	}
}