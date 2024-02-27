package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {

	url := flag.String("url", "https://www.google.com/", "url to parse")
	flag.Parse()

	if *url == "" {
		fmt.Fprintf(os.Stderr, "URL must be passed, ex : go run main.go --url http://google.com")
		os.Exit(1)
	}

	content, err := readHtml(*url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading html: %v\n", err)
		os.Exit(1)
	}

	doc, err := html.Parse(bytes.NewReader(content))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while parsing the html: %v\n", err)
		os.Exit(1)
	}

	var words, pics int
	var wordsMap = make(map[string]int)
	visit(doc, &words, &pics, &wordsMap)

	fmt.Printf("Words : %d, Pictures : %d\n", words, pics)
	for k, v := range wordsMap {
		fmt.Printf("Word : %s, Count : %d\n", k, v)
	}

}

func visit(node *html.Node, words, pics *int, wordsMap *map[string]int) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode && c.Parent.Data != "script" && c.Parent.Data != "style" {
			w := strings.Fields(c.Data)
			for _, word := range w {
				(*wordsMap)[word] = (*wordsMap)[word] + 1
			}
			*words += len(w)
		} else if c.Type == html.ElementNode {
			if c.Data == "img" {
				*pics++
			}
			visit(c, words, pics, wordsMap)
		}
	}

}

func readHtml(url string) ([]byte, error) {
	res, err := http.Get(url)

	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	html, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return html, nil

}
