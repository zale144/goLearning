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
		body = `<img src="https://upload.wikimedia.org/wikipedia/commons/0/06/
			Kitten_in_Rizal_Park%2C_Manila.jpg">`
	case "/dog":
		body = `<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/
			Golde33443.jpg">`
	default:
		body = "404 not found"
	}

	io.WriteString(res, body)
}

func main() {
	var h MyHandler

	http.ListenAndServe(":9000", h)
}
