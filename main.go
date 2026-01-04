package main

import (
	"net/http"

	"log"
	"os"
)


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/add",AddHandler)
	mux.HandleFunc("/sub",SubHandler)
	mux.HandleFunc("/mul",MulHandler)
	mux.HandleFunc("/div",DivHandler)
	Logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	Logger.Println("Server running on http://localhost:8080")
	Logger.Println("Try: http://localhost:8080/add?a=10&b=3")
	Logger.Println("Try: http://localhost:8080/sub?a=10&b=3")

	log.Fatal(http.ListenAndServe(":8080", mux))


}
