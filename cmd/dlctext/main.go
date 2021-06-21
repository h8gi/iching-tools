package main

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/antchfx/htmlquery"
)

const (
	baseUrl = "https://ctext.org/"
	menuUrl = "https://ctext.org/book-of-changes/yi-jing"
	outdir  = "./files"
)

var menu map[string]string

func main() {
	if _, err := os.Stat("html"); os.IsNotExist(err) {
		os.Mkdir(outdir, 0755)
	}

	doc, err := htmlquery.LoadURL(menuUrl)
	if err != nil {
		panic(err)
	}

	nodes, err := htmlquery.QueryAll(doc, "//div[@id='content3']/a")
	if err != nil {
		panic(err)
	}

	for _, n := range nodes {
		time.Sleep(1 * time.Second)

		href := htmlquery.SelectAttr(n, "href")
		url := baseUrl + href

		guadoc, err := htmlquery.LoadURL(url)
		if err != nil {
			panic(err)
		}
		table, err := htmlquery.Query(guadoc, "//table[2]")
		if err != nil {
			panic(err)
		}

		filename := outdir + "/" + path.Base(href) + ".html"

		fmt.Println("write file:", filename)
		f, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		f.Write([]byte(htmlquery.OutputHTML(table, true)))
	}

}
