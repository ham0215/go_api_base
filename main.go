package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ham0215/go_api_base/greetings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	message, err := greetings.Hello("hoge")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, message)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
