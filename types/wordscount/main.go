package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)

	scan.Split(bufio.ScanWords)

	words := make(map[string]int)
	totalwords := 0

	for scan.Scan() {
		words[scan.Text()]++
		totalwords++
	}

	fmt.Printf("The total word count is %v, and unique word count is %v\n", totalwords, len(words))

	type wordcount struct {
		word  string
		count int
	}

	var ss []wordcount

	for k, v := range words {
		ss = append(ss, wordcount{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].count > ss[j].count
		//return len(ss[i].word) > len(ss[j].word)
	})

	ss = ss[:3]
	for _, v := range ss {
		fmt.Printf("The word `%v` has count `%v`\n", v.word, v.count)
	}

	for _, v := range ss {
		fmt.Printf("The word `%v` has count `%v`\n", v.word, v.count)
	}

	//fmt.Println(ss)
}
