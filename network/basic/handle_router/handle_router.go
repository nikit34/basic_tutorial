package main

import (
	"fmt"
	"log"
	"net/http"
)


type logger struct {
	Inner http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Start")
	l.Inner.ServeHTTP(w, r)
	log.Println("Finish")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello\n")
}

func main() {
	f := http.HandlerFunc(hello)
	l := logger{
		Inner: f,
	}
	http.ListenAndServe(":8080", &l)
}