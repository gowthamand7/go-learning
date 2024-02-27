package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReadComic(t *testing.T) {
	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Respond with a sample comic JSON
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"month": "1",
			"num": 1,
			"link": "",
			"year": "2006",
			"news": "",
			"safe_title": "Barrel - Part 1",
			"transcript": "[[A boy sits in a barrel which is floating in an ocean.]]\nBoy: I wonder where I'll float next?",
			"alt": "Don't we all.",
			"img": "https://imgs.xkcd.com/comics/barrel_cropped_(1).jpg",
			"title": "Barrel - Part 1",
			"day": "1"
		}`))
	}))
	defer mockServer.Close()

	// Override the URL in the main package with the URL of the mock server
	url = mockServer.URL + "/1/info.0.json"

	var failcount int
	// Test ReadComic function with a valid comic index
	comic, err := ReadComic(url, &failcount)
	if err != nil {
		t.Errorf("ReadComic(1) returned error: %v", err)
	}
	if comic.Num != 1 {
		t.Errorf("Expected comic number to be 1, got %d", comic.Num)
	}
}

func TestReadComic404(t *testing.T) {
	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Respond with a sample comic JSON
		w.WriteHeader(http.StatusNotFound)
	}))
	defer mockServer.Close()

	// Override the URL in the main package with the URL of the mock server
	url = mockServer.URL + "/-1/info.0.json"

	var failcount int
	// Test ReadComic function with a invalid valid comic index
	comic, err := ReadComic(url, &failcount)
	if err == nil {
		t.Errorf("Expected error for invalid comic index, but got nil %v", comic)
	}
}

func TestReadComic500(t *testing.T) {
	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Respond with a sample comic JSON
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer mockServer.Close()

	// Override the URL in the main package with the URL of the mock server
	url = mockServer.URL + "/invalidid/info.0.json"

	var failcount int
	// Test ReadComic function with a invalid valid comic index
	comic, err := ReadComic(url, &failcount)
	if err == nil {
		t.Errorf("Expected error for invalid comic index, but got nil %v", comic)
	}
}

func TestReadComicFailedCount(t *testing.T) {
	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Respond with a sample comic JSON
		w.WriteHeader(http.StatusNotFound)
	}))

	// Override the URL in the main package with the URL of the mock server
	url = mockServer.URL + "/-1/info.0.json"

	var failcount int
	// Test ReadComic function with a valid comic index
	ReadComic(url, &failcount)

	if failcount != 1 {
		t.Error("Failcount should be 1, got ", failcount)
	}

	ReadComic(url, &failcount)

	if failcount != 2 {
		t.Error("Failcount should be 2, got ", failcount)
	}

	mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Respond with a sample comic JSON
		w.WriteHeader(http.StatusOK)
	}))

	defer mockServer.Close()

	ReadComic(url, &failcount)
	if failcount == 0 {
		t.Error("Failcount should be 2, got ", failcount)
	}

}
