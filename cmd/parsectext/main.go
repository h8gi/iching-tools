package main

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

type gua struct {
	id       string
	filepath string
	name     string
}

func main() {
	files, _ := filepath.Glob("./files/*.html")
	for _, fn := range files {
		base := path.Base(fn)
		id := strings.TrimSuffix(base, ".html")
		fmt.Println(id)
	}
}
