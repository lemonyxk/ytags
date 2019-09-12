package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ReadFileAsSlice(filePath string) (FileLine, FileByte) {

	f, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		log.Panicln(err)
	}
	defer func() { _ = f.Close() }()

	var reader = bufio.NewReader(f)

	var fileByte FileByte
	var fileLine FileLine
	var i = 0
	for {

		i++

		bytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fileByte = append(fileByte, bytes)
		fileLine = append(fileLine, i)
	}

	return fileLine, fileByte
}
