package main

import (
	"html/template"
	"log"
	"net/http"
	"github.com/zale144/goLearning/net/html/11-templates_csv/parse"
)

func main() {
	// parse csv
	records := parse.Parse("table.csv")

	// parse template
	tpl, err := template.ParseFiles("hw.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// function
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// execute template
		err = tpl.Execute(res, records)
		if err != nil {
			log.Fatalln(err)
		}
	})

	// create server
	http.ListenAndServe(":9000", nil)
}
