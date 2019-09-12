package main

import (
	"log"
	"os/exec"
)

func GoFmt(filePath string) {

	err := exec.Command("gofmt", "-w", filePath).Start()
	if err != nil {
		log.Panicln(err)
	}
}
