package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Trigram struct {
	Id            int64
	Unicode       int64
	Char          string
	Binary        uint
	Name          string
	Translation   string
	ImageInNature string
	Direction     string
}

func main() {
	// read trigram

	f, err := os.Open("./trigram.csv")

	if err != nil {
		panic(err)
	}

	r := csv.NewReader(f)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		tg := Trigram{
			Id:   record[0],
			Char: record[1],
		}

		for value := range record {
			fmt.Printf("%s\n", record[value])
		}
	}

}
