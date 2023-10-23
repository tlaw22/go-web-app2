package main

import (
	"fmt"
	"net/http"
)

// Home is the homepage handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 - %d", sum))
}

// AddValues adds 2 values togeter
func addValues(x, y int) int {
	return x + y
}
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/About", About)

	_ = http.ListenAndServe(":8080", nil)
}

// Naming a function with a lowercase start letter makes the function private
