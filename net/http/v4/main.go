package main

import (
	"net/http"
	"io"
)

type MyHandler int

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	var body string;
	switch req.URL.Path {
	case "/cat":
		body = "<strong>kitty kitty kitty</strong>"
	case "/dog":
		body = "<strong>doggy doggy doggy</strong>"
	default:
		body = "404 not found"
	}

	io.WriteString(res, body)
}

func main() {
	var h MyHandler

	http.ListenAndServe(":9000", h)
}
