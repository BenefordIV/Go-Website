package main

import (
	"net/http"
)

const portNumber = ":8080"

func main() {
	//fmt.Print("hello world")

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(portNumber, nil)
}
