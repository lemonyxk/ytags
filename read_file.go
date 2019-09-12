package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ReadFileAsSlice(file string) (FileLine, FileByte) {
	absFilePath, err := filepath.Abs(file)
	if err != nil {
		log.Panicln(err)
	}

	f, err := os.OpenFile(absFilePath, os.O_RDWR, 0666)
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
