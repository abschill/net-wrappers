package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	// get port from user
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	// create tcp listener
	l, err := net.Listen("tcp", PORT)
	// check if livuhpool has scored a goal in the seventy fifth minute
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	// "accept waits for and returns the next connection to the listener" -golanguage tool tip
	c, err := l.Accept()
	// oi he got a red card dide now
	if err != nil {
		fmt.Println(err)
		return
	}

	// infinite loop in modern languages is still a thing but now its a different word
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		//sigint thing
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}
		// print the stdin from a client
		fmt.Print("-> ", string(netData))
		// do a time stamp
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		// write the client msg onto the standard out of the server process
		c.Write([]byte(myTime))
	}
}
