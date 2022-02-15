package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Trigram struct {
	Id                 int
	Unicode            rune
	Char               string
	Binary             uint64
	BinaryString       string
	Name               string
	Translation        string
	ImageInNature      string
	Direction          string
	FamilyRelationship string
	BodyPart           string
	Attribute          string
	Stage              string
	Animal             string
}

func main() {
	// read trigram

	f, err := os.Open("./trigram.csv")

	if err != nil {
		panic(err)
	}

	r := csv.NewReader(f)

	tgs := make([]Trigram, 0, 8)

	r.Read()

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		id, _ := strconv.Atoi(record[0])
		bin, _ := strconv.ParseUint(record[2], 2, 64)

		tg := Trigram{
			Id:                 id,
			Unicode:            []rune(record[1])[0],
			Char:               record[1],
			Binary:             bin,
			BinaryString:       record[2],
			Name:               record[3],
			Translation:        record[4],
			ImageInNature:      record[5],
			Direction:          record[6],
			FamilyRelationship: record[7],
			BodyPart:           record[8],
			Attribute:          record[9],
			Stage:              record[10],
			Animal:             record[11],
		}

		tgs = append(tgs, tg)
	}

	trijson, _ := json.MarshalIndent(tgs, "", "	")

	fmt.Print(string(trijson))

}
