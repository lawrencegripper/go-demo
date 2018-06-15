package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8087", nil)
	fmt.Println(err)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
