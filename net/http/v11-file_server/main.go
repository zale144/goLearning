package main

import (
	"net/http"
	"strings"
	"io"
)

func upTown(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	var dogName string
	fs := strings.Split(req.URL.Path, "/")
	if len(fs) >= 3 {
		dogName = fs[2]
	}
	// the image doesn't serve
	io.WriteString(res, `
		Dog Name: <strong>` + dogName + `</strong><br/>
		<img src="/toby.jpg">
		<img src="/geisha.jpg">
	`)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./img")))
	http.HandleFunc("/dog/", upTown)
	http.ListenAndServe(":9000", nil)
}
