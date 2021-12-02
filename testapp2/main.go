package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
	log.Info("handler called")
}

func main() {
	log.SetLevel(log.InfoLevel)

	file, err := os.OpenFile("/logs/testapp2.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Info("Failed to log to file, using default stderr")
	} else {
		log.SetOutput(file)
		defer file.Close()
	}

	port := os.Getenv("PORT")

	fmt.Printf("testapp2 started at port %s\n", port)
	log.Infof("testapp2 started at port %s", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
