package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	files, _ := filepath.Glob("./files/*.html")
	fmt.Println(files)
}
