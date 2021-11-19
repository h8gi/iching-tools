package main

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/antchfx/htmlquery"
)

type gua struct {
	id       string
	filepath string
	name     string
	data     content
}

type content struct {
	TuanZhuan  string `json:"彖傳"`
	XiangZhuan string `json:"象傳"`
}

func main() {
	// var guatable map[string]gua

	files, _ := filepath.Glob("./files/*.html")
	for _, fn := range files {
		base := path.Base(fn)
		id := strings.TrimSuffix(base, ".html")
		fmt.Println(id)
		parseFile(fn)
	}
}

func parseFile(filepath string) error {
	doc, err := htmlquery.LoadDoc(filepath)
	if err != nil {
		return err
	}
	clist := htmlquery.Find(doc, "//td[contains(@class, 'ctext')]")
	elist := htmlquery.Find(doc, "//td[contains(@class, 'etext')]")
	for i := range clist {
		fmt.Println(htmlquery.InnerText(clist[i]))
		fmt.Println(htmlquery.InnerText(elist[i]))
	}
	return nil
}
