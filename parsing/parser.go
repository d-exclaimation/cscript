//
//  parser.go
//  parsing
//
//  Created by d-exclaimation on 1:45 pm.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package parsing

import (
	"fmt"
	"os"
	"strings"
)

type Parser struct {
	content string
	module  string
	headers []string
	sources []string
}

func NewParser(filename string) (*Parser, error) {
	module, _, found := strings.Cut(filename, ".c")
	if !found {
        module, _, found = strings.Cut(filename, ".h")
        if !found {
		    return nil, fmt.Errorf("filename of %s is not a C source code", filename)
        }
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &Parser{
		content: string(content),
		module:  module,
		headers: make([]string, 0),
		sources: make([]string, 0),
	}, nil
}

func (p *Parser) CompileWithFlag() {
	isBeforeFlag := true
	lines := strings.Split(p.content, "\n")
	for i, line := range lines {
		if i+1 < len(lines) && IsFunctionDeclaration(line, lines[i+1]) {
			isBeforeFlag = false
			p.headers = append(p.headers, "", fmt.Sprintf("%s;", line))
		}

		if isBeforeFlag || IsDefinition(line) {
			p.headers = append(p.headers, line)
		} else {
			p.sources = append(p.sources, line)
		}
	}
}

func (p *Parser) Generate() error {
	if err := os.WriteFile(p.module+".c", []byte(p.sourceContent()), 0644); err != nil {
		return err
	}
	f, err := os.Create(p.module + ".h")
	if err != nil {
		return err
	}
	_, err = f.WriteString(p.headerContent())
	if err != nil {
		return err
	}
	return nil
}
