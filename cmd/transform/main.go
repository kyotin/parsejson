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
	inJson    = flag.String("inJson", "/Users/tinnguyen/Downloads/test.json", "path to json file")
	out       = flag.String("out", "./out.json", "path to out template file")
	workers   = flag.String("workers", "1", "max number of workers")
	buffLines = flag.String("buffLines", "100", "buffer lines when reading")
)

type _Source struct {
	PersonName                   string `json:"person_name"`
	PersonFirstNameUnanalyzed    string `json:"person_first_name_unanalyzed"`
	PersonLastNameUnanalyzed     string `json:"person_last_name_unanalyzed"`
	PersonNameUnanalyzedDowncase string `json:"person_name_unanalyzed_downcase"`
	PersonEmailStatusCd          string `json:"person_email_status_cd"`
	PersonExtrapolatedEmail      string `json:"person_extrapolated_email"`
	PersonEmail                  string `json:"person_email"`
	PersonLinkedinUrl            string `json:"person_linkedin_url"`
	SantizedOrganizationName     string `json:"sanitized_organization_name_unanalyzed"`
	OrganizationName             string `json:"organization_name"`
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

	buffLines, _ := strconv.Atoi(*buffLines)
	lines := make(chan string, buffLines)
	go func() {
		numOfLines := 0

		// Start reading from the file with a reader.
		reader := bufio.NewReader(in)
		for {
			var buffer bytes.Buffer
			endOfFile := false
			for {
				l, isPrefix, err := reader.ReadLine()
				if err == io.EOF {
					endOfFile = true
					break
				}

				buffer.Write(l)
				// If we've reached the end of the line, stop reading.
				if !isPrefix {
					break
				}

				if err != nil && err != io.EOF {
					fmt.Printf("ERROR %v \n", err)
					break
				}
			}

			if endOfFile {
				break
			}

			line := buffer.String()
			if line != "" {
				lines <- line
				numOfLines += 1
			}
		}

		fmt.Printf("Sending %d lines \n", numOfLines)
		close(lines)
	}()

	maxWorker, _ := strconv.Atoi(*workers)

	goodLines := make(chan string, maxWorker)
	go func(goodLines <-chan string) {
		for line := range goodLines {
			if _, err := out.WriteString(line + "\n"); err != nil {
				fmt.Errorf("Can't write string to file")
			}
		}
	}(goodLines)

	var wg sync.WaitGroup
	for i := 0; i < maxWorker; i++ {
		wg.Add(1)
		go func(workerId int, lines <-chan string, goodLines chan<- string, wg *sync.WaitGroup) {
			fmt.Printf("Worker %d Start \n", workerId)
			numOfLines := 0
			for line := range lines {
				numOfLines += 1
				record := &Record{}
				err := json.Unmarshal([]byte(line), record)
				if err != nil {
					fmt.Printf("Can't parse json from line: %s \n", line)
				} else {
					// DO business here
					if b, err := json.Marshal(record.Source); err == nil {
						goodLines <- string(b)
					} else {
						fmt.Println(err)
					}
				}
			}

			fmt.Printf("Worker %d had procesed %d lines \n", workerId, numOfLines)
			wg.Done()
		}(i, lines, goodLines, &wg)
	}

	wg.Wait()
	close(goodLines)
}
