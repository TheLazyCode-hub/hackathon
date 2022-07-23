package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World - root context</h1>")
}

func internal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World - internal context</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/internal", internal)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)
}
