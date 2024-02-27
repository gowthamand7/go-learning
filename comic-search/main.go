package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	op := flag.String("type", "", "Type of operation to perform")

	flag.Parse()

	if *op == "" || *op != "search" && *op != "load" {
		fmt.Fprintln(os.Stderr, "Error: type of operation must be specified using the -type flag allowed values are: search, load")
		os.Exit(1)
	}

	if *op == "load" {
		Load()
	}

	if *op == "search" {

		if len(flag.Args()[1:]) < 1 {
			fmt.Fprintln(os.Stderr, "Error: No search terms provided")
			os.Exit(1)
		}

		fn := flag.Arg(0)
		if fn == "" {
			fmt.Fprintln(os.Stderr, "Error: No input file provided")
			os.Exit(1)
		}

		var comics Comics

		json.Unmarshal(SafeReadFile(fn), &comics)

		c := Search(flag.Args()[1:], comics)
		fmt.Fprintf(os.Stdout, "Found comic: %+v\n", c)
	}

}

func SafeReadFile(fn string) []byte {

	var (
		f   *os.File
		err error
		c   []byte
	)

	if f, err = os.Open(fn); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input file:", err)
		os.Exit(1)
	}

	defer f.Close()

	if c, err = io.ReadAll(f); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input file:", err)
		os.Exit(1)
	}

	return c
}
