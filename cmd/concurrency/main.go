package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"github.com/sirupsen/logrus"
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

	lines := make(chan string, 2000)
	go func() {
		scanner := bufio.NewScanner(in)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		close(lines)
	}()

	goodLines := make(chan string, 100)
	go func(goodLines <-chan string) {
		for line := range goodLines {
			if _, err := out.WriteString(line + "\n"); err != nil {
				logrus.Errorf("Can't write string to file")
			}

		}
	}(goodLines)

	maxWorker := 100
	for i:=0; i < maxWorker; i++ {
		go func(lines <-chan string, goodLines chan<- string){
			for line := range lines {
				record := &Record{}
				err := json.Unmarshal([]byte(line), record)
				if err != nil {
					logrus.Errorf("Can't parse json.", line)
				} else {
					// DO business
					if record.Source.PersonPhone != "" {
						goodLines <- line
					}
				}
			}
		}(lines, goodLines)
	}

}
