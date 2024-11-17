package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! This message is generated using GO without dockerfile")
}

func main() {
	http.HandleFunc("/", handler)
	port := ":8080"
	fmt.Printf("Server running on port %s...\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
