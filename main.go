package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	serverPort = ":3000"
)

func main() {
	fmt.Printf("----------------------\n")
	fmt.Printf("# Fútbol Backend\n")
	fmt.Printf("----------------------\n")
	fmt.Printf("# Server listening to: http://localhost%s\n", serverPort)

	log.Fatal(http.ListenAndServe(serverPort, GetRouter()))
}

func logger(request *http.Request) {
	fmt.Fprintf(os.Stdout, "[%s] %s\n", request.Method, request.RequestURI)
}
