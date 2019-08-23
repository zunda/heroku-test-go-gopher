package main

import (
	"bytes"
	"fmt"
	"io"
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
		log.Printf("Request from %s", r.RemoteAddr)
		fmt.Fprintf(w, "Go Gopher!\n")
		switch r.Method {
		case "GET":
		case "HEAD":
		case "POST":
			buf := &bytes.Buffer{}
			size, err := io.Copy(buf, r.Body)
			if err != nil {
				w.WriteHeader(503)
				log.Fatal(err)
			} else {
				fmt.Fprintf(w, "Request body was %d bytes\n", size)
			}
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
