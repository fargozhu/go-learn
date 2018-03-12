package main

import (
	"fmt"
	"log"
	"net/http"
)

type Timezone int

const (
	EST Timezone = -(5 + iota)
	CST          //CST -6
	MST          //MST -7
	PST          //PST -8
)

func main() {
	http.HandleFunc("/hello", handHello)
	fmt.Println("serving on http://0.0.0.0:8080/hello")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func handHello(w http.ResponseWriter, req *http.Request) {
	log.Println("serving %s in EST%d timezone", req.URL, EST)
	fmt.Fprintln(w, "Hello")
}
