package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
	fmt.Println("handler called")
}

func main() {
	port := os.Getenv("PORT")

	fmt.Printf("testapp1 started at port %s\n", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
