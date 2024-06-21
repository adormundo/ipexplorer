package main

import (
	"fmt"
	"net/http"
)

func main() {
	ipAddr := getLocalIP()
	http.HandleFunc("/", handleIPRequest)
	addr := ":8080"
	fmt.Printf("Server started at http://%s%s\n", ipAddr, addr)
	http.ListenAndServe(addr, nil)
}