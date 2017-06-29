package main

import (
	"os"
	"fmt"
	"log"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		log.Println("err happened", err)
		//log.Fatalln(err)
		//panic(err)
		print("here")
	}
}
