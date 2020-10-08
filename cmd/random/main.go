package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var (
	inJson        = flag.String("inJson", "/Users/tinnguyen/Downloads/test.json", "path to json file")
	out           = flag.String("out", "./out.json", "path to out template file")
	numOfRows     = flag.Int("numOfRows", 10000, "number of rows wanna extract")
	maxRowsInFile = flag.Int("maxRowsInFile", 1000000, "max rows in file")
)

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

	// generate number
	pickedRow := make(map[int]struct{}, *numOfRows)

	rand.Seed(time.Now().UnixNano())
	for {
		min := 1
		max := *maxRowsInFile
		pickedRow[rand.Intn(max-min+1) + min] = struct{}{}

		if len(pickedRow) == *numOfRows {
			break
		}
	}

	// Start reading from the file with a reader.
	numOfLines := 0
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
			numOfLines += 1
			if  _, ok := pickedRow[numOfLines]; ok {
				if _, err := out.WriteString(line + "\n"); err != nil {
					fmt.Println("Can't write string to file")
				}
			}
		}
	}

}
