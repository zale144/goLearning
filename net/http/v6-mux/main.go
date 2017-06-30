package main

import (
	"net/http"
	"io"
)

type dogHandler int
type catHandler int
type ratHandler int

func (h dogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `<img src="https://i.ytimg.com/vi/opKg3fyqWt4/hqdefault.jpg" >`)
}

func (h catHandler) ServeHTTP(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `<img src="http://cdn3.teen.com/wp-content/uploads/2015/09/kitten-flower-crown-main.jpg" >`)
}

func (h ratHandler) ServeHTTP(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `<img src="https://media.mnn.com/assets/images/2014/01/rat-wearing-knit-hat-JF.jpg.838x0_q80.jpg" >`)
}

func main() {
	var dog dogHandler
	var cat catHandler
	var rat ratHandler

	mux := http.NewServeMux()
	mux.Handle("/dog/", dog)
	mux.Handle("/cat/", cat)
	mux.Handle("/", rat)

	http.ListenAndServe(":9000", mux)
}
