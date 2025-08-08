package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	dumpRequest := false
	if os.Getenv("DUMP_REQUEST") != "" {
		dumpRequest = true
	}
	closeConnection := false
	if os.Getenv("CONNECTION_CLOSE") != "" {
		closeConnection = true
	}

	h := http.NewServeMux()
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request from %s to %s", r.RemoteAddr, r.URL.Path)
		if dumpRequest {
			dump, err := httputil.DumpRequest(r, false)
			if err != nil {
				http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
				return
			}
			log.Printf("%q", dump)
		}

		if closeConnection {
			w.Header().Set("Connection", "close")
		}

		stat := r.Header.Get("X-Response-Status")
		if stat != "" {
			s, err := strconv.Atoi(stat)
			if err == nil {
				log.Println("Responding with HTTP status code: " + stat)
				w.WriteHeader(s)
			}
		}

		switch r.Method {
		case "GET":
			switch r.URL.Path {
			case "/echo":
				fmt.Fprintf(w, "%s %s %s\nHost: %s\n", r.Method, r.URL.Path, r.Proto, r.Host)
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
			case "/favicon.ico":
				http.ServeFile(w, r, "favicon.ico")
			case "/ref.png":
				http.ServeFile(w, r, "ref.png")
			default:
				fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
<head><title>Go Gopher!</title></head>
<body><p>Go Gopher!</p><p><img src="/ref.png"></p></body>
</html>
`)
			}
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

	log.Println("Listening at port " + port)
	err := http.ListenAndServe(":"+port, h)
	log.Fatal(err)
}
