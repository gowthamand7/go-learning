package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World " + r.URL.Path))
		w.Header().Set("Status", "400 OK")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
