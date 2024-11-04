package main

import (
	"net/http"

	"github.com/Beneford_IV/udemy_app/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	//fmt.Print("hello world")

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_ = http.ListenAndServe(portNumber, nil)
}
