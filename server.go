package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "hello\n")
}

func headers(writer http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(writer, "%v: %v\n", name, h)
		}
	}
}

func main() {
	log.Print("starting server...")
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting port to %s", port)
	}

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
