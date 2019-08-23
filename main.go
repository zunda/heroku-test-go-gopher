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
		log.Printf("Request from %s", r.RemoteAddr);
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "Go Gopher!\n")
		case "HEAD":
			fmt.Fprintf(w, "Go Gopher!\n")
		default:
			w.WriteHeader(405)
		}
	})
	h.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "favicon.ico")
	})

	log.Println("Listening at port " + port)
	err := http.ListenAndServe(":"+port, h)
	log.Fatal(err)
}
