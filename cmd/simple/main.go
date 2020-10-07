package main

import (
	"encoding/json"
	"flag"
	"io"
	"os"
)

var (
	inJson = flag.String("inJson", "", "path to json file")
	out    = flag.String("out", "", "path to out template file")
)

type _Source struct {
	PersonName  string `json:"person_name"`
	PersonEmail string `json:"person_email"`
	PersonPhone string `json:"person_phone"`
}

type Record struct {
	Index  string  `json:"_index"`
	Type   string  `json:"_type"`
	Source _Source `json:"_source"`
}


func main() {
	flag.Parse()

	in, err := os.Open(*inJson)
	if err != nil {
		panic("Can't open file")
	}
	defer in.Close()

	out, err := os.Create(*out)
	if err != nil {
		panic("Can't create file")
	}
	defer out.Close()

	dec := json.NewDecoder(in)
	for {
		record := &Record{}
		if err := dec.Decode(record); err == io.EOF {
			break
		} else if err != nil {
			continue
		}

		// business
		if record.Source.PersonPhone != "" {

		}
	}

}
