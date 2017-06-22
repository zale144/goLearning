package main

import (
	"fmt"
	//"Tutorial/01_helloWorld/utils"
)


var b string = "neki string iz maina"
var c = 144

func main() {
	mapa := map[int]string{
		0 : "aleks",
		1 : "jen",
		2 : "antonina",
	}
	for k, v := range mapa {
		fmt.Println(k, v)
	}
}

