package main

import (
	"fmt"
	"net/http"
)

func main() {
	//fmt.Print("hello world")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "hello world")
		if err != nil {
			fmt.Printf("error found %s", err.Error())
		}

		fmt.Printf("bytes written: %d", n)
	})

	_ = http.ListenAndServe(":8080", nil)
}
