package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, server minimal go: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":3333", nil)
}