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
	input   string
	output  string
	headers []string
	sources []string
}

func NewParser(input string, output string) (*Parser, error) {
	inputName, err := module(input)
	if err != nil {
		return nil, fmt.Errorf("filename of %s is not a C source code", input)
	}
	content, err := os.ReadFile(input)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &Parser{
		content: string(content),
		input:   inputName,
		output:  output,
		headers: make([]string, 0),
		sources: make([]string, 0),
	}, nil
}

func (p *Parser) CompileWithFlag() {
	isBeforeFlag := true
	lines := strings.Split(p.content, "\n")
	for i, line := range lines {
		if IsStaticDeclaration(line) || (i+1 < len(lines) && IsFunctionDeclaration(line, lines[i+1])) {
			isBeforeFlag = false
			if !IsStaticDeclaration(line) {
				p.headers = append(p.headers, fmt.Sprintf("%s;", line), "")
			}
		}

		if isBeforeFlag || IsDefinition(line) {
			p.headers = append(p.headers, line)
		} else {
			p.sources = append(p.sources, line)
		}
	}
}

func (p *Parser) Generate() error {
	if err := os.WriteFile(p.output+".c", []byte(p.sourceContent()), 0644); err != nil {
		return err
	}
	if err := os.WriteFile(p.output+".h", []byte(p.headerContent()), 0644); err != nil {
		return err
	}
	return nil
}
