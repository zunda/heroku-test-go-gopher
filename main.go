package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	h := http.NewServeMux()
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go Gopher!\n")
	})
	h.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "favicon.ico")
	})

	log.Println("Listening at port " + port)
	err := http.ListenAndServe(":"+port, h)
	log.Fatal(err)
}
