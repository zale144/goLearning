package main

import (
	"net/http"
	"io"
)

type MyHandler int

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request)  {
	var body string;
	switch req.URL.Path {
	case "/cat":
		body = "kitty kitty kitty"
	case "/dog":
		body = "doggy doggy doggy"
	default:
		body = "404 not found"
	}

	io.WriteString(res, body)
}

func main() {
	var h MyHandler

	http.ListenAndServe(":9000", h)
}
