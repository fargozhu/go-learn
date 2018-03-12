package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handHello)
	fmt.Println("serving on http://0.0.0.0:8080/hello")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func handHello(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.URL)
	fmt.Fprintln(w, "Hello")
}
