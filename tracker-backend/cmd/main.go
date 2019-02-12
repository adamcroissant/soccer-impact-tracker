package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/api/v1/games", GetGames)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
