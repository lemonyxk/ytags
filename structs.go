package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func Do(path string, tags []string) {

	log.Println(path, tags)

	var fset = token.NewFileSet()

	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(node ast.Node) bool {

		n, ok := node.(*ast.StructType)
		if !ok {
			return true
		}

		for i := 0; i < len(n.Fields.List); i++ {

			// if n.Fields.List[i].Tag == nil {
			// 	continue
			// }

			if len(n.Fields.List[i].Names) == 0 {
				continue
			}

			var causeKey = n.Fields.List[i].Names[0].Name

			if n.Fields.List[i].Tag == nil {
				n.Fields.List[i].Tag = &ast.BasicLit{}
			}
			n.Fields.List[i].Tag.Value = CreateTags(causeKey, tags)
		}

		return true
	})

	var output []byte
	buffer := bytes.NewBuffer(output)
	err = format.Node(buffer, fset, f)
	if err != nil {
		panic(err)
	}

	write(path, buffer)

}

func write(path string, buf *bytes.Buffer) {

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}

	defer func() { _ = f.Close() }()

	_, err = f.WriteString(buf.String())
	if err != nil {
		panic(err)
	}

}

func CreateTags(causeKey string, tags []string) string {

	// `json:"name"`

	var createTags = "`"
	for _, v := range tags {
		createTags += fmt.Sprintf("%s:\"%s\" ", v, CauseWord(causeKey))
	}
	createTags = createTags[:len(createTags)-1]
	createTags += "`"

	return createTags
}

func IsBigWord(char byte) bool {
	return char >= 65 && char <= 90
}

func CauseWord(word string) string {
	if !*fix {
		return word
	}
	var s bytes.Buffer
	for i, v := range word {
		if i > 0 && IsBigWord(byte(v)) && ((i+1 < len(word) && !IsBigWord(word[i+1])) || !IsBigWord(word[i-1])) {
			s.WriteString("_")
		}
		if IsBigWord(byte(v)) {
			s.WriteString(string(v + 32))
		} else {
			s.WriteString(string(v))
		}
	}
	return s.String()
}
