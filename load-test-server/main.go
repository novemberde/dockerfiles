package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func load() {
	a := 0
	for i := 0; i < 3000000; i++ {
		a = a*a + a*2 + i
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	load()

	fmt.Fprintf(w, "Hi there, your url path is %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	var port string

	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
