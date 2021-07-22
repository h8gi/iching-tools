package main

import "github.com/antchfx/htmlquery"

const url = "https://en.wikipedia.org/wiki/List_of_hexagrams_of_the_I_Ching"

func main() {
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		panic(err)
	}

}
