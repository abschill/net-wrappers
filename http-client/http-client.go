package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	// make sure the http-server is running before trying
	// call the main url
	resp, err := http.Get("http://127.0.0.1")
	if err != nil {
		panic(err)
	}

	// print status of response
	println(fmt.Sprintf("Status: %d", resp.StatusCode))

	if resp.StatusCode == 200 {
		scanner := bufio.NewScanner(resp.Body)
		for i := 0; scanner.Scan(); i++ {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	} else {
		panic("Bad Status")
	}
}
