package main

import (
	"fmt"

	"github.com/antchfx/htmlquery"
)

const path = "./List_of_hexagrams_of_the_I_Ching.html"

func main() {
	doc, err := htmlquery.LoadDoc(path)
	if err != nil {
		panic(err)
	}

	titles, err := htmlquery.QueryAll(doc, "//h2/span[starts-with(@id, 'Hexagram')]")

	texts, err := htmlquery.QueryAll(doc, "//h2/span[starts-with(@id, 'Hexagram')]/../following-sibling::p[1]")

	tables, err := htmlquery.QueryAll(doc, "//h2/span[starts-with(@id, 'Hexagram')]/../following-sibling::table[1]")
	if err != nil {
		panic(err)
	}

	for i := range titles {
		if err != nil {
			panic(err)
		}
		title := htmlquery.InnerText(titles[i])
		text := htmlquery.InnerText(texts[i])
		table := htmlquery.InnerText(tables[i])
		fmt.Println(title, text, table)
	}
}
