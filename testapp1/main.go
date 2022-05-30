package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
	fmt.Println("handler called")
}

func main() {
	port := os.Getenv("PORT")

	fmt.Printf("testapp1 started at port %s\n", port)

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		fmt.Println(errors.Unwrap(err))
	}
}
