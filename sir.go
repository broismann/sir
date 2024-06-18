package main

import (
	"fmt"
	"net/http"
	"os"
	"sir/handler"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./sir <port>")
		os.Exit(1)
	}
	portStr := os.Args[1]
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("Invalid port number:", err)
		os.Exit(1)
	}

	{
		http.HandleFunc("/", handler.MyHandler())
		fmt.Println("Starting HTTP server on port", port)
		err := http.ListenAndServe(":0"+strconv.Itoa(port), nil)
		if err != nil {
			fmt.Println("Error starting HTTP server")
		}
	}
}
