package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	_ = os.Setenv("TZ", "Asia/Shanghai")
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
}

var fix = flag.Bool("x", false, "fixed the rule")
var tags = flag.String("t", "", "tags")
var file = flag.String("f", "", "file")

func main() {

	flag.Parse()

	var tagsSlice = strings.Split(*tags, ",")
	var filePath = *file

	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		println(err)
		return
	}

	f, err := os.Stat(absFilePath)
	if os.IsNotExist(err) {
		println(absFilePath, "not exists")
		return
	}

	if f.IsDir() {
		println(absFilePath, "is a dir")
		return
	}

	Do(absFilePath, tagsSlice)

	println("tags rebuild success")
}
