package main

import (
	"flag"
)

func main() {
	var from, to string
	var limit, offset int

	flag.StringVar(&from, "from", "", "/path/to/file")
	flag.StringVar(&to, "to", "", "/path/to/file")
	flag.IntVar(&limit, "limit", 0, "1024")
	flag.IntVar(&offset, "offset", 0, "1024")

	flag.Parse()
}

func Copy(src, dst string, limit, offset int) error {

}