package main 

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, `{"message": "Hello from Go Backend"}`)
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}