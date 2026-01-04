package main

import (
	"fmt"
	"net/http"
	"strconv"
	"log"
	"os"
)
var logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
func AddHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := ValidateParam(w, r)
	logger.Printf("Performing addition")
	if err != nil {
		return
	}
	result := add(a, b)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Result : %.2f + %.2f = %.2f", a, b, result)
	logger.Printf("Status: %d,Ip Address %s, Path: %s",http.StatusOK,r.RemoteAddr,r.URL.Path)
}

func SubHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := ValidateParam(w, r)
	if err != nil {
		return
	}
	result := sub(a, b)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Result : %.2f - %.2f = %.2f", a, b, result)

	logger.Printf("Status: %d,Ip Address %s, Path: %s",http.StatusOK,r.RemoteAddr,r.URL.Path)
}

func MulHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := ValidateParam(w, r)
	if err != nil {
		return
	}
	result := multi(a, b)
	w.Header().Set("Content-Type", "text/plain")

	logger.Printf("Status: %d,Ip Address %s, Path: %s",http.StatusOK,r.RemoteAddr,r.URL.Path)
	fmt.Fprintf(w, "Result : %.2f * %.2f = %.2f", a, b, result)
}

func DivHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := ValidateParam(w, r)
	if err != nil {
		return
	}
	if b == 0 {
		http.Error(w, "value of 'b' cannot be 0", http.StatusBadRequest)
		return
	}
	result := div(a, b) // Fixed: was calling multi instead of div
	w.Header().Set("Content-Type", "text/plain")

	logger.Printf("Status: %d,Ip Address %s, Path: %s",http.StatusOK,r.RemoteAddr,r.URL.Path)
	fmt.Fprintf(w, "Result : %.2f / %.2f = %.2f", a, b, result)
}

func ValidateParam(w http.ResponseWriter, r *http.Request) (float64, float64, error) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	if aStr == "" || bStr == "" {
		http.Error(w, "Both 'a' and 'b' are required", http.StatusBadRequest)
		return 0, 0, fmt.Errorf("Missing Param")
	}

	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		http.Error(w, "param 'a' must be a valid number", http.StatusBadRequest)
		return 0, 0, fmt.Errorf("'a' is not a valid number")
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		http.Error(w, "param 'b' must be a valid number", http.StatusBadRequest)
		return 0, 0, fmt.Errorf("'b' is not a valid number")
	}

	return a, b, nil
}
