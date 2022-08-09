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
