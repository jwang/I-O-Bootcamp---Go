package main

import (
    "fmt"
    "http"
)

type String string
func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, s)
}

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}
func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, s)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, ", r.FormValue("name"))
}

func main() {
    http.HandleFunc("/", handler)
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", Struct{"Hello", ":", "Gophers!"})
    http.ListenAndServe(":8080", nil)
}