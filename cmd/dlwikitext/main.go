package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
)

const path = "./List_of_hexagrams_of_the_I_Ching.html"

type Hexagram struct {
	Id      int64
	Unicode int64
	Char    string
	Text    string
}

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

	gualist := make([]Hexagram, len(titles), len(titles))

	for i := range titles {
		if err != nil {
			panic(err)
		}
		title := htmlquery.InnerText(titles[i])

		id, err := strconv.ParseInt(strings.TrimLeft(title, "Hexagram "), 10, 64)
		if err != nil {
			panic(err)
		}

		text := htmlquery.InnerText(texts[i])

		unicode, err := htmlquery.Query(tables[i], "//td/a[@title='Unicode']/../following-sibling::td[1]")
		if err != nil {
			panic(err)
		}

		unicodeInt, err := strconv.ParseInt(htmlquery.InnerText(unicode), 10, 64)
		if err != nil {
			panic(err)
		}

		unicodeChar := string(unicodeInt)

		gua := Hexagram{
			Id:      id,
			Unicode: unicodeInt,
			Char:    unicodeChar,
			Text:    text,
		}

		gualist[i] = gua
	}

	guajson, err := json.MarshalIndent(gualist, "", "	")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(guajson))
}
