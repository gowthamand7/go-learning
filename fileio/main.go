package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	/**
	-l : only print no of lines in the given file(s)
	-w : only print no of words in the given file(s)
	-c : only print no of characters in the given file(s)
	-m : print the size of the given file(s)
	**/
	var options Options
	options.l = flag.Bool("l", false, "only print no of lines in the given file(s)")
	options.w = flag.Bool("w", false, "only print no of words in the given file(s)")
	options.c = flag.Bool("c", false, "only print no of characters in the given file(s)")
	options.m = flag.Bool("m", false, "displays count of characters from a file(s)")
	options.L = flag.Bool("L", false, "Only Displays maximum line length of a file(s)")
	flag.Parse()

	var globalfilesstat FileStat = FileStat{
		filename: "Total",
	}

	filenames := flag.Args()

	for _, filename := range filenames {
		f, err := os.Open(filename)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		var filestat = FileStat{
			filename: filename,
		}

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			filestat.lines++
			text := scanner.Text()
			filestat.words += len(strings.Fields(text))
			if len(text) > filestat.maxlinelength {
				filestat.maxlinelength = len(text)
			}
		}
		fileinfo, _ := os.Stat(filename)
		filestat.chars += int(fileinfo.Size())

		globalfilesstat.lines += filestat.lines
		globalfilesstat.words += filestat.words
		globalfilesstat.chars += filestat.chars
		if globalfilesstat.maxlinelength < filestat.maxlinelength {
			globalfilesstat.maxlinelength = filestat.maxlinelength
		}

		f.Close()
		print(filestat, options)
	}

	if len(filenames) > 1 {
		print(globalfilesstat, options)
	}

}

func print(fs FileStat, o Options) {

	switch {
	case *o.l:
		fmt.Printf("%7d %7s\n", fs.lines, fs.filename)
	case *o.w:
		fmt.Printf("%7d %7s\n", fs.words, fs.filename)
	case *o.c:
		fmt.Printf("%7d %7s\n", fs.chars, fs.filename)
	case *o.m:
		fmt.Printf("%7d %7s\n", fs.chars, fs.filename)
	case *o.L:
		fmt.Printf("%7d %7s\n", fs.maxlinelength, fs.filename)
	default:
		fmt.Printf("%7d %7d %7d %7s\n", fs.lines, fs.words, fs.chars, fs.filename)
	}

}

type FileStat struct {
	filename      string
	lines         int
	words         int
	chars         int
	maxlinelength int
}

type Options struct {
	l *bool
	w *bool
	c *bool
	m *bool
	L *bool
}
