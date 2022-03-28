package main

import (
	"encoding/json"
	"flag"
	"fmt"
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

	file, err := os.Open(*fileName)

	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	fc := FileContent{}
	err = json.NewDecoder(file).Decode(&fc)
	if err != nil {
		log.Panicln(err)
	}

	for _, p := range fc.People {
		fmt.Println(p.Firstname)
	}
}
