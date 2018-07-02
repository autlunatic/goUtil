package main

import (
	"fmt"

	"github.com/autlunatic/goUtil/Zipping"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("hell")
	toFileName := "yolo.zip"
	fromFileNames := []string{"main.go"}

	err := Zipping.ZipFiles(toFileName, fromFileNames)
	if err != nil {
		color.Red(fmt.Sprint("ERROR! > ", err))
	} else {
		fmt.Println("done")
	}
}
