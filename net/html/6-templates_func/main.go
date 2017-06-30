package main

import (
	"text/template"
	"log"
	"os"
)

type Page struct {
	Title string
	Body string
}

func main() {
	var err error

	tpl := template.New("tpl.gohtml") // pointer to a template
	tpl = tpl.Funcs(template.FuncMap{
		"mycustomfunc": func() string {
			return "This should work"
		},
	})

	tpl, err = tpl.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, Page{
		Title: "My Title",
	})
	if err != nil {
		log.Fatalln(err)
	}
}