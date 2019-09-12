package main

import (
	"log"
	"os/exec"
	"path/filepath"
)

func GoFmt(file string) {
	absFilePath, err := filepath.Abs(file)
	if err != nil {
		log.Panicln(err)
	}
	err = exec.Command("gofmt", "-w", absFilePath).Start()
	if err != nil {
		log.Panicln(err)
	}
}
