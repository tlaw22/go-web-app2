package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the homepage handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)

	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 = %d", sum))
}

// AddValues adds 2 values togeter
func addValues(x, y int) int {
	return x + y
}
func Divvide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
}
func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by Zero")
		return 0, err
	}
	result := x / y
	return result, nil
}
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/About", About)
	http.HandleFunc("/divide", Divvide)
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	fmt.Println()
	_ = http.ListenAndServe(portNumber, nil)
}

// Naming a function with a lowercase start letter makes the function private
