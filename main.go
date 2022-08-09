//
// main.go
// cscript
//
// Created by d-exclaimation on 00:00.
//

package main

import (
	"github.com/d-exclaimation/cscript/parsing"
	"log"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	p, err := parsing.NewParser(filename)
	if err != nil {
		log.Fatalln(err)
	}

	p.CompileWithFlag()
	if err := p.Generate(); err != nil {
		log.Fatalln(err)
	}
}

func filenames() (string, string) {
	input := os.Args[1]
	output := ""
	if len(os.Args) < 3 {
		output = strings.Split(input, ".")[0]
	} else {
		output = os.Args[2]
	}
	return input, output
}
