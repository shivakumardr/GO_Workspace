package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a Home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("This is a About and sum of 2 + 2 is:%d ", sum))
}

// Divide is the dived  page handler
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divieValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Can't divide by zero")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f id divide by %f is %f ", 100.0, 10.0, f))
}

// addValues adds two int and return sum of int
func addValues(x, y int) int {
	return x + y
}

func divieValues(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("Con't divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting the application at port: %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
