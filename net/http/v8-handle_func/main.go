package main

import (
	"net/http"
	"io"
)

func upTown(res http.ResponseWriter, req *http.Request)  {
	io.WriteString(res,"doggy doggy doggy")
}

func youUp(res http.ResponseWriter, req *http.Request)  {
	io.WriteString(res,"catty catty catty")
}

func main() {

	http.HandleFunc("/", upTown)
	http.HandleFunc("/cat/", youUp)

	http.ListenAndServe(":9000", nil) // pass nil to use DefaultServeMux ( = NewServeMux() )
}

// we use http.HandleFunc instead of *ServeMux.HandleFunc