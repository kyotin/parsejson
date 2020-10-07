package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
)

var (
	inJson = flag.String("inJson", "/Users/tinnguyen/Downloads/test.json", "path to json file")
	out    = flag.String("out", "./out.json", "path to out template file")
	workers = flag.String("worker", "100", "max number of workers")
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
		numOfLines := 0

		// Start reading from the file with a reader.
		reader := bufio.NewReader(in)
		for {
			var buffer bytes.Buffer
			for {
				l, isPrefix, err := reader.ReadLine()
				buffer.Write(l)
				// If we've reached the end of the line, stop reading.
				if !isPrefix {
					break
				}
				// If we're at the EOF, break.
				if err != nil {
					if err != io.EOF {
						fmt.Printf("ERROR %v \n", err)
					}
					break
				}
			}
			line := buffer.String()
			if line != "" {
				lines <- line
			}

			if err == io.EOF {
				break
			}

			numOfLines += 1
		}

		if err != io.EOF {
			fmt.Printf("Can't process whole file - Failed with error: %v\n", err)
		}

		fmt.Println("Sending %d lines", numOfLines)
		close(lines)
	}()

	goodLines := make(chan string, 100)
	go func(goodLines <-chan string) {
		for line := range goodLines {
			if _, err := out.WriteString(line + "\n"); err != nil {
				fmt.Errorf("Can't write string to file")
			}
		}
	}(goodLines)

	var wg sync.WaitGroup

	maxWorker, _ := strconv.Atoi(*workers)
	for i:=0; i < maxWorker; i++ {
		wg.Add(1)
		go func(workerId int, lines <-chan string, goodLines chan<- string, wg *sync.WaitGroup){
			numOfLines := 0
			for line := range lines {
				numOfLines += 1
				record := &Record{}
				err := json.Unmarshal([]byte(line), record)
				if err != nil {
					fmt.Printf("Can't parse json from line: %s \n", line)
				} else {
					// DO business
					if record.Source.PersonPhone != "" {
						goodLines <- line
					}
				}
			}
			defer func() {
				fmt.Printf("Worker %d had procesed %d lines \n", workerId, numOfLines)
				wg.Done()
			}()
		}(i, lines, goodLines, &wg)
	}

	wg.Wait()
}
