package main

import (
	"fmt"

	"github.com/antchfx/htmlquery"
)

const url = "https://en.wikipedia.org/wiki/List_of_hexagrams_of_the_I_Ching"

const path = "./List_of_hexagrams_of_the_I_Ching.html"

func main() {
	doc, err := htmlquery.LoadDoc(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(doc.FirstChild)
}
