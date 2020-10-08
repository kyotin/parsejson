package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"hash/fnv"
	"io"
	"os"
)

var (
	inJson        = flag.String("inJson", "/Users/tinnguyen/Downloads/test.json", "path to json file")
	out           = flag.String("out", "./out.json", "path to out template file")
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

	hashedLines := make(map[uint32]struct{}, *maxRowsInFile)
	duplicated := 0

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
			if _, ok := hashedLines[hash(line)]; ok {
				duplicated += 1
			} else {
				hashedLines[hash(line)] = struct{}{}
				out.WriteString(line + "\n")
			}
		}
	}

	fmt.Printf("Duplicated: %d \n", duplicated)
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
