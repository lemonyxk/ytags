package main

type FileByte [][]byte
type FileLine []int

type ChangeObject struct {
	ChangeByte FileByte
	ChangeLine FileLine
	StructName string
	Comment    string
}
