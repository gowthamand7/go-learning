package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	url        = "https://xkcd.com/%d/info.0.json"
	startindex = 1
)

func Load() {
	// Buffer to store JSON data
	var b bytes.Buffer
	defer b.Reset()

	// Slice to store fetched comics
	var comics []Comic

	// Variable to track consecutive failed attempts
	var failcount int

	// Loop until two consecutive fetch failures occur
	for ; failcount < 2; startindex++ {
		// Construct the URL for the current comic
		localurl := fmt.Sprintf(url, startindex)
		// Read the comic from the URL
		comic, err := ReadComic(localurl, &failcount)
		if err != nil {
			// Print error and continue to the next comic
			fmt.Fprintf(os.Stderr, "%s \n", err.Error())
			continue
		}
		// Append the fetched comic to the comics slice
		comics = append(comics, comic)
	}

	// Marshal comics slice into JSON format
	j, err := json.Marshal(Comics{comics})
	if err != nil {
		// Print error if JSON marshaling fails
		fmt.Fprintf(os.Stderr, "error while marshaling the comics : %v \n", err)
		os.Exit(1)
	}

	// Write JSON data to the buffer and print it
	fmt.Fprintln(&b, string(j))
	fmt.Println(b.String())
}

// ReadComic fetches comic data from the provided URL
func ReadComic(url string, failcount *int) (Comic, error) {
	// Send GET request to fetch comic data
	resp, err := http.Get(url)
	if err != nil {
		// Return error if fetching fails
		return Comic{}, fmt.Errorf("error while fetching comic %s : %v", url, err.Error())
	}

	// Increment failcount if comic is not found
	if resp.StatusCode == http.StatusNotFound {
		*failcount++
	} else {
		// Reset failcount if comic is found
		*failcount = 0
	}

	// Return error if response status code is not OK
	if resp.StatusCode != http.StatusOK {
		return Comic{}, fmt.Errorf("error while fetching comic %s with the %d code, url %s", url, resp.StatusCode, url)
	}

	// Close response body when done
	defer resp.Body.Close()

	// Read response body
	j, _ := io.ReadAll(resp.Body)

	// Unmarshal JSON data into Comic struct
	var comic Comic
	err = json.Unmarshal(j, &comic)
	if err != nil {
		// Return error if JSON unmarshaling fails
		return Comic{}, fmt.Errorf("error while unmarshal the json for comic %s : %v", url, err.Error())
	}

	return comic, nil
}
