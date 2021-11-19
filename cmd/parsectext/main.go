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
	list := htmlquery.Find(doc, "//td[contains(@class, 'ctext')]")
	for _, l := range list {
		fmt.Println(htmlquery.InnerText(l))
	}
	return nil
}
