package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received!")
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8087", nil))
}
