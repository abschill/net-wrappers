package main

import (
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("simple text server demo.\n"))
}

func main() {
	http.HandleFunc("/", Hello)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
