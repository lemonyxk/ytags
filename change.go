package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Change(filePath string, changeObjectSlice []*ChangeObject, fileLine FileLine, fileByte FileByte, tags []string) {
	
	absFilePath, err := filepath.Abs(filePath)
	println(absFilePath)
	if err != nil {
		log.Panicln(err)
	}
	
	for _, object := range changeObjectSlice {
		
		for index, value := range object.ChangeByte {
			var line, content = object.ChangeLine[index], string(value)
			var changeContent = Do(line, content, tags)
			if changeContent == "" {
				continue
			}
			fileByte[line-1] = []byte(changeContent)
		}
	}
	
	f, err := os.OpenFile(absFilePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Panicln(err)
	}
	defer func() { _ = f.Close() }()
	
	for _, value := range fileByte {
		_, err := f.Write(value)
		_, err = f.Write([]byte("\r\n"))
		if err != nil {
			log.Panicln(err)
		}
	}
	
}

func Do(line int, content string, tags []string) string {
	var contentSlice []string
	
	var temp = strings.Split(strings.TrimSpace(content), " ")
	
	for _, value := range temp {
		var s = strings.TrimSpace(value)
		if s == "" {
			continue
		}
		contentSlice = append(contentSlice, strings.TrimSpace(value))
	}
	
	if len(contentSlice) == 0 || len(contentSlice) == 1 {
		return ""
	}
	
	var key = contentSlice[0]
	
	if !IsBigWord(key[0]) {
		return ""
	}
	
	var causeKey = CauseWord(key)
	
	var createTags = CreateTags(causeKey, tags)
	
	if len(contentSlice) > 2 {
		contentSlice = contentSlice[0:3]
		contentSlice[2] = createTags
	} else {
		contentSlice = append(contentSlice, createTags)
	}
	
	return strings.Join(contentSlice, " ")
}

func CreateTags(causeKey string, tags []string) string {
	
	// `json:"name"`
	
	var createTags = "`"
	for _, v := range tags {
		createTags += fmt.Sprintf("%s:\"%s\" ", v, causeKey)
	}
	createTags += "`"
	
	return createTags
}

func IsBigWord(char byte) bool {
	return char >= 65 && char <= 90
}

func CauseWord(word string) string {
	var s = ""
	for i, v := range word {
		if i > 0 && IsBigWord(byte(v)) && ((i+1 < len(word) && !IsBigWord(word[i+1])) || !IsBigWord(word[i-1])) {
			s += "_"
		}
		if IsBigWord(byte(v)) {
			s += string(v + 32)
		} else {
			s += string(v)
		}
	}
	return s
}
