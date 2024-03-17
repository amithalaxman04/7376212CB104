package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")

	})
	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Amitha")

	})
	http.HandleFunc("/register_number", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376212cb104")

	})
	http.HandleFunc("/department", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "csbs")

	})
	http.HandleFunc("/fav_colour", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "black")

	})

	fmt.Printf("Server running (port=8080), route: http://localhost:8080/helloworld\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
