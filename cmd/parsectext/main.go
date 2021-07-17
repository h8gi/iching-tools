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
	彖傳 string
	象傳 string
}

func main() {
	var guatable map[string]gua

	files, _ := filepath.Glob("./files/*.html")
	for _, fn := range files {
		base := path.Base(fn)
		id := strings.TrimSuffix(base, ".html")
		fmt.Println(id)
	}
}

func parseFile(filepath string) (content, error) {
	doc, err := htmlquery.LoadDoc(filepath)
	if err != nil {
		return gua{}, err
	}

}
