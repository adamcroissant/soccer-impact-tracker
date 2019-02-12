package main

import (
	"net/http"
	"fmt"
	"log"
	"html"
)

func main() {
	http.HandleFunc("/api/v1/games", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}