package main

import (
	"strings"
)

func Search(fileLine FileLine, fileByte FileByte) []*ChangeObject {

	var changeObjectSlice []*ChangeObject

	var changeObject *ChangeObject

	var shouldAdd = false

	for line, bytes := range fileByte {

		var content = string(bytes)

		if strings.TrimSpace(content) == "" {
			continue
		}

		if strings.Index(content, "struct {") > -1 {
			shouldAdd = true

			changeObject = &ChangeObject{}
			changeObject.StructName = strings.Split(content, " ")[1]

			if strings.Index(string(fileByte[line-1]), "//") > -1 {
				changeObject.Comment = strings.TrimSpace(string(fileByte[line-1][2:]))
			}
			continue
		}

		if string(bytes[0]) == "}" && shouldAdd == true {
			shouldAdd = false
			changeObjectSlice = append(changeObjectSlice, changeObject)
			changeObject = nil
			continue
		}

		if shouldAdd == true && changeObject != nil {
			changeObject.ChangeByte = append(changeObject.ChangeByte, bytes)
			changeObject.ChangeLine = append(changeObject.ChangeLine, fileLine[line])
		}

	}

	return changeObjectSlice

}
