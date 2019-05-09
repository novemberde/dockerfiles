package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, your url path is %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)

	if port := os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
