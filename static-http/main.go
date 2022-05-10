package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func serveHTML(w http.ResponseWriter, req *http.Request) {
	// get file as []byte
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		panic(err)
	}
	// give the byte array to the writer
	w.WriteHeader(200)
	w.Write(content)
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":80", nil)
}
