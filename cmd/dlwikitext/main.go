package main

import (
	"fmt"
	"strconv"

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

		unicode, err := htmlquery.Query(tables[i], "//td/a[@title='Unicode']/../following-sibling::td[1]")
		if err != nil {
			panic(err)
		}

		i, err := strconv.ParseInt(htmlquery.InnerText(unicode), 10, 64)
		if err != nil {
			panic(err)
		}

		hexarune := string(i)

		fmt.Println(title, text, hexarune)
	}
}
