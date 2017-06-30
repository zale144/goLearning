package main

import (
	"text/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, 10) // the "." is the current data/pipeline
	if err != nil {
		log.Fatalln(err)
	}
}