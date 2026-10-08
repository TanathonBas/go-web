package main

import (
	"net/http"
	"fmt"
)

func helloGOHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go!")
}

func hiGOHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, Go!")
}

func main() {
	http.HandleFunc("/hellogo", helloGOHandler)
	http.HandleFunc("/higo", hiGOHandler)

	fmt.Printf("Starting server on :3030....")

	if err := http.ListenAndServe(":3030", nil); err != nil {
		fmt.Printf("Error starting server: %v", err)
	}

	
}