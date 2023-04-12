package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	_ = os.Setenv("TZ", "Asia/Shanghai")
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
}

var fix = HasArgs([]string{"-x"})
var tags = GetArgs([]string{"-t"})
var file = GetArgs([]string{"-f"})

func main() {

	var tagsSlice = strings.Split(tags, " ")
	var filePath = file

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

func GetArgs(flag []string) string {
	args := os.Args
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(flag); j++ {
			if args[i] == flag[j] {
				if i+1 < len(args) {
					return args[i+1]
				}
			}
		}
	}
	return ""
}

func HasArgs(flag []string) bool {
	args := os.Args
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(flag); j++ {
			if args[i] == flag[j] {
				return true
			}
		}
	}
	return false
}
