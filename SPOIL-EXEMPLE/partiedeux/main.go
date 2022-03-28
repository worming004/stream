package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var fileName = flag.String("f", "", "filename to parse")

type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type FileContent struct {
	People []Person `json:"people"`
}

func main() {
	flag.Parse()

	var reader io.Reader
	if *fileName != "" {

		file, err := os.Open(*fileName)

		if err != nil {
			log.Panicln(err)
		}
		defer file.Close()

		reader = file
	} else {
		reader = os.Stdin
	}

	fc := FileContent{}
	err := json.NewDecoder(reader).Decode(&fc)
	if err != nil {
		log.Panicln(err)
	}

	for _, p := range fc.People {
		fmt.Println(p.Firstname)
	}
}
