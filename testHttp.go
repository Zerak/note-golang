package main

import (
	"log"
	"net/http"
)

type String string

func (s String)ServeHTTP(rw http.ResponseWriter, req *http.Request){
	log.Printf("%v\n",req.URL.Path)
	rw.Write(([]byte)("hello String" ))
}


type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s Struct)ServeHTTP(rw http.ResponseWriter, req *http.Request){
	log.Printf("%v\n",req.URL.Path)
	rw.Write(([]byte)("hello Struct "))
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})

	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}