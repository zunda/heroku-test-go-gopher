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
		log.Printf("Request from %s to %s", r.RemoteAddr, r.URL.Path)

		switch r.URL.Path {
		case "/echo":
			for name, values := range r.Header {
				for _, value := range values {
					fmt.Fprintf(w, "%s: %s\n", name, value)
				}
			}
		case "/session":
			sid := "unknown session ID"
			c, err := r.Cookie("heroku-session-affinity")
			if err == nil {
				sid = c.Value
			}
			dyno := os.Getenv("DYNO")
			if dyno == "" {
				dyno = "unknown dyno"
			}
			fmt.Fprintf(w, "Hello Gopher from %s with %s!\n", dyno, sid)
		default:
			fmt.Fprintf(w, "Go Gopher!\n")
		}

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
