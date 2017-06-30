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
		<img src="/resources/toby.jpg">
		<img src="/resources/geisha.jpg">
	`)
}

func main() {
	http.Handle("/resources/",
		// so it doesn't look for "./assets/resources/toby.jpg"
		http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/dog/", upTown)
	http.ListenAndServe(":9000", nil)
}
