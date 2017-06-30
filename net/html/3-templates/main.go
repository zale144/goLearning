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
	err = tpl.Execute(os.Stdout, []int{1,2,3,4,5,}) // the "." is the current data/pipeline
	if err != nil {
		log.Fatalln(err)
	}
}