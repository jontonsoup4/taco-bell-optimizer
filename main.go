package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":8000"
	router := NewRouter()
	fmt.Printf("Serving tacos on %s\r\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
