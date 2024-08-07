package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server on :8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("aaaaaaaaaaaaaaaaa")
		fmt.Fprintf(w, "Hello, World!aaaaaaaaaa")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
