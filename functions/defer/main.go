package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	flag.Parse()
	args := flag.Args()

	fmt.Println(do())

	j := new(int)

	fmt.Println(*j)

	fmt.Println(*returnAddr())
	i := 0
	for _, filename := range args {
		size, err := readfile(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer fmt.Printf("%6d %6s\n", size, filename)

		i := 5
		fmt.Printf(" inner i stored @ %p", &i)
	}
	fmt.Printf("outter i stored @ %p", &i)
}

func readfile(filename string) (size int64, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	stat, _ := f.Stat()
	size = stat.Size()
	return
}

func returnAddr() *int {
	i := 10
	return &i
}

func do() int {
	return 10
}
