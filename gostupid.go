package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Starting up gostupid...\n")

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/ws", socketHandler)
	http.ListenAndServe(":8000", nil)
}
